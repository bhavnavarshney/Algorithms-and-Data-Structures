class Node{
	int a ;
	Node next;
	
	Node(int x,Node z){
		a=x;
		next = z;
	}
	
}

class LinkedList{
	Node head;
	Node append(int data){
		Node n = new Node(data,null);
		if(head==null){
			head = n;
		}
		else{
			Node temp = head;
			while(temp.next!=null){
				temp = temp.next;
				
			}
			temp.next = n;
		}
		return head;
	}
	
	Node appendStart(int data,Node head){
	
		Node head2 = new Node(data,head);
		return head2;
	}
	
	void print(Node head){
		Node temp = head;
		while(temp!=null){
			System.out.println(temp.a+" ");
			temp = temp.next;
		}
	}
}

class LinkedListImplementation{
	public static void main(String s[]){
		LinkedList ll = new LinkedList();
		Node head = ll.append(10);
		head = ll.append(20);
		head = ll.append(30);
		head = ll.append(40);
		head = ll.append(50);
		
		ll.print(head);
		
		Node head2 = ll.appendStart(0,head);
		
		ll.print(head2);
	}
}