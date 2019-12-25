import java.io.FileInputStream;
import java.io.IOException;
import java.nio.charset.StandardCharsets;

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