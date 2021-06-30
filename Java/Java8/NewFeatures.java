package Java8;

public class NewFeatures {
    public static void main(String[] args) {
//        Interfaces8 i1 = new Interfaces8();
//        i1.act();
//        defaultInterfaceEg i2 = new defaultInterfaceEg();
//        i2.act();

        ForEachLoop fe = new ForEachLoop();
        fe.whatIsForEachLoop();

        LambdaFeature lf = new LambdaFeature();
        lf.whatIsLamdaFeature();

        StreamsDemo sd = new StreamsDemo();
        sd.whatIsStream();

        MethodRef mr = new MethodRef();
        mr.howIsMethodPassedAsRef();
    }

}
