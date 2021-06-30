package Java8;

import java.util.Arrays;
import java.util.List;
import java.util.Locale;

interface Parser {
    String parse(String s);
}

class StringParser {
    public static String convert(String s) {
        if(s.length()<=3) {
            s = s.toUpperCase();
        }
        return s;
    }
}
class MyClass {
    public void print(String s, Parser p) {
        s = p.parse(s);
        System.out.println(s);
    }
}
public class MethodRef {
    public void howIsMethodPassedAsRef() {
        List<Integer> l = Arrays.asList(1,2,3,4);
        l.forEach(System.out::println);

        //Eg: 2
        MyClass c = new MyClass();
        c.print("Bhavna", StringParser::convert);
    }
}
