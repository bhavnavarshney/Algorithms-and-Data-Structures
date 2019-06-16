import java.util.*;
class Maximum{
	static Scanner sc = new Scanner(System.in);
	
	public static void main(String s[]){
		
		int n = sc.nextInt();
		int a[] = new int[n];
		for(int i=0;i<n;i++){
			a[i] = sc.nextInt();
		}
		int sum = find(n,a);
		System.out.println(sum);
	}
	public static int  find(int n, int[] a){
		int sum = a[0];
		for(int i=1;i<a.length;i++){
			if(a[i] > sum)
				sum = a[i];
		}
		return sum;
	}
} 
