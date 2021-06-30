package Java8;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class StreamsDemo {
    void whatIsStream() {
        // It is used for processing the data. Can be used on once!

        List<Integer> l = new ArrayList<>();
        for(int i=1; i<=100; i++) {
            l.add(i);
        }
        System.out.println(l.stream().filter(i->{
            if(i==2) return true;
            return false;
        }).findAny().orElse(0));
    }
}
