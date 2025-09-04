class Node {
    constructor(data) {
        this.data = data;
        this.next = null
    };
}

export const inputMaker = () => {
    let head = new Node(1);
    head.next = new Node(2);
    head.next.next = new Node(3);
    head.next.next.next = new Node(4);
    head.next.next.next.next = new Node(5);
    head.next.next.next.next.next = new Node(6);

    return head;
}

export const printList = (curr) => {
    let output = "";
    while (curr !== null) {
        output += curr.data + " ";
        curr = curr.next;
    }

    console.log(output.trim());
}