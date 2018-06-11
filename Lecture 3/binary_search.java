// Binary Search

class binary_search{
	public static void main(String s[]){
		int a[] = {1,5,7,8,13,15,17};
		int x = 8;
		
		int start = 0,end = 7;
		
		for(start=0;start<end;start++){
			int mid = (start+end)/2;
		
			if(a[mid]==x){
				System.out.println("Found");
				break;
			}
			if(a[mid]>x){
				end = mid;
			}
			if(a[mid]<x){
				start = mid;
			}
		
		}
}
}