class inheritance{
	public static void main (String s[]){
		Child c = new Child();
		Parent  p = new Child();
		//Child c2 = new Parent();
		//c2.method1();
		//c.method1();
		p.method1();
	}
}
class  Parent{
	void method1(){
		System.out.println("Hey! I'm in parent");
	}
}

class Child extends Parent{
	void method1(){
		System.out.println("Hey! im in child");
	}
}