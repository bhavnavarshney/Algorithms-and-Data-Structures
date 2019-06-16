// Sum of rows,cols,diagonal 

class 2D{
	public static void main(String s[]){
		
		int a[][] = new int[5][5];
		for(int i=0;i<n;i++){
			for(int j=0;j<n;j++){
				a[i][j] = i*j+1;
			}
		}
		
		// Sum of rows 
		int rsum=0,csum=0,dsum=0;
		for(int i=0;i<5;i++){
			rsum=0;
			for(int j=0;j<n;j++){
				rsum += a[i][j];
			}
			System.out.println("Sum of row "+i+" = "+rsum);
		}
		
		// Sum of cols
		
		
		
		// Sum of diagonal
		
		
	}
}