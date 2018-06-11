import java.util.*;
class Selection{
	public static void main(String s[]){
		Scanner sc = new Scanner(System.in);
		int n = sc.nextInt();
		int a[] = new int[n];
		for(int i=0;i<n;i++)
			a[i] = sc.nextInt();
		
		int i=0,max,t=0,current=0;
		//for(i=0;i<n;i++){
			for(int j=i+1;j<n;j++){
				max = findMax(i,a);
				t = a[max];
				a[max] = a[i];
				a[i] = t;
				i++;
			}
			//a[i] = max;
		//}
		
		for( i=0;i<n;i++)
			System.out.print(a[i]+" ");
	}
	
	public static int findMax(int current,int a[]){
		for(int i=current+1;i<a.length;i++){
			if(a[i]>a[current])
				current = i;
		}
		return current;
	}
}