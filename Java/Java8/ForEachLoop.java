package Java8;
import java.util.Arrays;
import java.util.List;

public class ForEachLoop {

    public void whatIsForEachLoop() {
        List<Integer> arrayList = Arrays.asList(1,2,3,4,5);
        arrayList.forEach(i -> System.out.println(i)); //Lambda
    }
}
