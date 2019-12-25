# Deploying a properties file next to your jar on OpenShift

I've recently came to the situation where i deployed a spring boot application on openshift and a certain dependency needed a properties file. The
 only problem was that this dependency didn't scan the classpath for the file, but just opened a FileInputStream relative to your current path.
 
 In this blogpost i will guide you through the process of deploying a text file next to a jar in an openshift container. I've reproduced an
  example scenario, but you can skip these steps and go right to the solution if you want to.
  
  ## The Example Scenario
  To reproduce the problem just described we need a jar with some java code reading a file located next to the jar. 
  
  #### The code:
  
  ```
public class ReadFileExample {

    public static void main(String[] args) throws IOException {
        String propsLocation = System.getenv("PROPS_LOCATION");
        try (FileInputStream fis = new FileInputStream(propsLocation)) {
            int i = 0;
            do {

                byte[] buf = new byte[1024];
                i = fis.read(buf);

                String value = new String(buf, StandardCharsets.UTF_8);
                System.out.print(value);

            } while (i != -1);
        }
    }
}
```

The code just opens a FileInputStream which location is taken from an environment property. 
The contents of the file are then printed in the outputstream.

#### The input file
I've created a properties.example file containing: `welcome bob`

If you want to run the code locally you'll need to set the PROPS_LOCATION env var:

`export PROPS_LOCATION=../resources/properties.example`

#### Creating a jar
First of we need to compile the code:

`javac ReadFileExample.java -d ../output/`

then add a MANIFEST.MF with the main class to run, in the same output dir:

```
Manifest-Version: 1.0
Main-Class: ReadFileExample
```

and then to create the jar:

` jar cfm readfileexample.jar MANIFEST.MF *.class`

Place the jar in a specific dir, we'll need it when creating the image on openshift.

### Creating and deploying an image from the jar
To create and deploy an image we need to specify to openshift the objects we need.
We need a BuildConfig, ImageStream and DeploymentConfig

i've specified the needed objects in yaml format in an openshift template:

```
apiVersion: template.openshift.io/v1
kind: Template
metadata:
  name: readfileexample-template
  annotations:
    openshift.io/display-name: readfileexample-template
labels:
  template: readfileexample-template

objects:
  - apiVersion: image.openshift.io/v1
    kind: ImageStream
    metadata:
      name: readfile-example
      labels:
        app: readfile-example

  - kind: BuildConfig
    apiVersion: v1
    metadata:
      name: readfile-example
    spec:
      strategy:
        sourceStrategy:
          from:
            kind: ImageStreamTag
            name: java:8
            namespace: openshift
        type: Source
      source:
        type: Binary
        binary:
          asFile: readfileexample.jar
      output:
        to:
          kind: ImageStreamTag
          name: readfile-example:latest

  - apiVersion: apps.openshift.io/v1
    kind: DeploymentConfig
    metadata:
      name: readfile-example
    spec:
      replicas: 1
      selector:
        app: readfile-example
      strategy:
        type: Rolling
      test: false
      template:
        metadata:
          labels:
            app: readfile-example
        spec:
          containers:
            - command:
                - "sh"
              args:
                - "-c"
                - "while true; do echo hello; sleep 18000;done"
              image: " "
              imagePullPolicy: IfNotPresent
              name: readfile-example
      triggers:
        - imageChangeParams:
            automatic: true
            containerNames:
              - readfile-example
            from:
              kind: ImageStreamTag
              name: readfile-example:latest
          type: ImageChange
        - type: ConfigChange

```

In the container spec i've added a shell command to execute to keep the container alive.
It's a simple infinite loop so that the container doesn't become idle and won't get killed by openshift.

### Creating and mounting the input file on openshift
To create and mount an input file we need to add a ConfigMap to our template:

```
  - apiVersion: v1
    kind: ConfigMap
    metadata:
      name: properties-configmap
    data:
      properties.example: welcome bob
```

And change our deploymentconfig to add a volume with our configmap, and mount this volume add a specific path

```
     template:
        metadata:
          labels:
            app: readfile-example
        spec:
          volumes:
            - name: properties-volume
              configMap:
                name: properties-configmap
          containers:
            - volumeMounts:
                - mountPath: "deployments/config"
                  name: properties-volume
```

All we need now is to add an environment variable to our deployment pointing to the mounted file from our configmap

```
              env:
                - name: PROPS_LOCATION
                  value: "config/properties.example"
```


### Running the jar
In the openshift console go to your deployment and click on the created pod and enter the terminal,
here you'll find the jar and mounted input file in the `/deployments` dir.

here is the result:
```
sh-4.2$ java -jar readfileexample.jar 
welcome bob
```
