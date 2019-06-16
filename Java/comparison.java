class  comparison{
	
	public static void main(String s[]){
		String s1 = new String("DEF");
		String s2 = new String("DEF");
		String s3 = "xyz";
		String s4 = "xyz";
		System.out.println(s3==s4);
		System.out.println(s3.equals(s4));
		System.out.println(s1==s2);
		System.out.println(s1.equals(s2));
}
}