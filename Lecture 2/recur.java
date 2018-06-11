import java.util.*;
class recur{
	public static void main(String s[]){
		Scanner sc = new Scanner(System.in);
		int no = sc.nextInt();
		System.out.println(fact(1,no));
		System.out.println(iter(no));
	}
	public static int fact(int n,int no){
		if(n==no)
			return n;
		else
			return n*fact(n+1,no);
	}
	public static int iter(int no){
		int fact=1;
		for(int i=2;i<=no;i++){
			fact = fact *i;
		}
		return fact;
	}
}