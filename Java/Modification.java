class Modification{
	public static void main(String s[]){
		/*int a = 5;
		modify(a);
		System.out.println(a);
		*/
		
		Balloon b = new Balloon("red");
		modify(b);
		System.out.println(b.color);
	}
	
	public static void modify(Balloon c){
		c.color="grey";
		//c = new Balloon("grey");
	}
}

class Balloon{
	
	String color="blue";
	Balloon(String y){
		color=y;
	}
}