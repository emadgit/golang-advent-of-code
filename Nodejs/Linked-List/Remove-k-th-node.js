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

import { inputMaker, printList } from "./helpers.js";

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