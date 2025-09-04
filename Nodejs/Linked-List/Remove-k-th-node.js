/*

Given a singly linked list, the task is to remove every kth node of the linked list.
Assume that k is always less than or equal to the length of the Linked List.

Examples :

Input: LinkedList: 1 -> 2 -> 3 -> 4 -> 5 -> 6, k = 2
Output: 1 -> 3 -> 5
Explanation: After removing every 2nd node of the linked list, the resultant linked list will be: 1 -> 3 -> 5 .

Input: LinkedList: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10, k = 3
Output: 1 -> 2 -> 4 -> 5 -> 7 -> 8 -> 10
Explanation: After removing every 3rd node of the linked list, the resultant linked list will be: 1 -> 2 -> 4 -> 5 -> 7 -> 8 -> 10.

*/

class Node {
    constructor(data) {
        this.data = data;
        this.next = null
    };
}

const inputMaker = () => {
    let head = new Node(1);
    head.next = new Node(2);
    head.next.next = new Node(3);
    head.next.next.next = new Node(4);
    head.next.next.next.next = new Node(5);
    head.next.next.next.next.next = new Node(6);

    return head;
}

const printList = (curr) => {
    let output = "";
    while (curr !== null) {
        output += curr.data + " ";
        curr = curr.next;
    }

    console.log(output.trim());
}

const solution = (k) => {
    let head = inputMaker();

    if (head === null || k<= 0) {
        return head;
    }

    let currentNode = head;
    let prev = null;
    let count = 0;

    while (currentNode != null) {
        count ++;

        if (count % k === 0) {
            if (prev !== null) {
                prev.next = currentNode.next;
            } else {
                head = currentNode.next;
            }
        } else {
            prev = currentNode
        }
        currentNode = currentNode.next;
    }
    printList(head)
}

solution(2);