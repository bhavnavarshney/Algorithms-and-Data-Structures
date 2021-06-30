package Java8;

interface MyInterface {
    public void something(String s);
}

//class A implements MyInterface {
//    public void something() {
//        System.out.println("Hi, there is a better way of implementing me! So comment me");
//    }
//}
public class LambdaFeature {
    public void whatIsLamdaFeature() {
//        MyInterface a = new MyInterface() {
//            public void something(String s) {
//                System.out.println("So much code!" + s);
//            }
//        }

        // Let's do the same above implementation using lamda!
        MyInterface a ;
        // a = (String s) -> System.out.println("Hey! This is lamda function and much better way of implementing me! If you want to see the magic, checkout what is created after you compile me!");
        // As there is only 1 parameter - You can write like below without braces and type!
        // This is implemented using Consumer interface!
        a = s-> System.out.println("Hey! This is lamda function and much better way of implementing me! If you want to see the magic, checkout what is created after you compile me, "+s);

        a.something("Developer");
    }
}
