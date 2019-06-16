import java.util.*;
class memorization{
	public static void main(String s[]){
		Scanner sc = new Scanner(System.in);
		int test = sc.nextInt();
		int input[] = new int[test];
		
		for(int i=0;i<test;i++){
			input[i] = sc.nextInt();
		}
		int max = Maximum.find(test,input);
		
		System.out.println("Maximum no = "+max);
		
		int f[] = new int[test+1];
		f = factorial(max);

		for(int i=0;i<test;i++){
			System.out.println(f[input[i]]);
		}
		
	}
	public static int[] factorial(int n){
		int factorials[] = new int[n+1];
		factorials[0] = 1;
		int fact=1;
		for(int i=1;i<=n;i++){
			
			factorials[i] =( i * factorials[i-1])%10000007;
		}
		return factorials;
	} 
}