import java.util.*;
class array{
	public static void main(String s[]){
		int[] a = new int[100];
		modify(a);
		Scanner sc = new Scanner(System.in);
		System.out.println(sc);
		System.out.println(a[0]); //Passed value of reference; 
	}
	public static void modify(int[] a){
		a[0] = 10;
	}
}