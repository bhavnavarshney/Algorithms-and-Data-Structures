class Student{
	
	static int no_of_student=0;
	
	String name;
	int grade=0; 
	int marks;
	
	/*void get(String a , String b , int c){
		name = a;
		grade = b;
		marks = c;
	} */
	
	Student(){
		no_of_student++;
	}
	
	void print(){
		System.out.println("Grade = "+grade);
	} 
	
	void clearSemester(){
		grade += 1;
	}
	
	static void decrementCount(){
		no_of_student--;
	}
	
}

class StudentDemo{
	public static void main(String s[]){
		Student s1 = new Student();
		Student s2 = new Student();
		Student s3 = new Student();
		Student s4 = new Student();
		Student s5 = new Student();
		
		s1.clearSemester();
		s1.print();
		s2.clearSemester();
		s2.print();
		s3.clearSemester();
		s3.print();
		s4.clearSemester();
		s4.print();
		s5.clearSemester();
		s5.print();
		
		s1.clearSemester();
		s1.print();
		
		System.out.println("No of Student = "+Student.no_of_student);
		Student.decrementCount();
		System.out.println("No of Student = "+Student.no_of_student);
		
	}
}