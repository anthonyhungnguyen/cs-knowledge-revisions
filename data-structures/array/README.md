# Array

#compsci #data-structures

## Basic

-   The linear data structure **collects elements of the same data type and stores them in contiguous and adjacent memory locations.**
-   Declaration: int arr[5]
-   **Indexing**: elements in an array are accessed using an index. The first element has an index of 0, and the last has an index of n-1
-   **Traversing**: Loops
-   **Operations**: search, sort, insert, delete
-   **Memory allocation**: Allocated in contiguous blocks, they use a fixed amount of memory. This means it is **efficient in memory usage**, but the size of an array cannot be changed at runtime.

![Image](https://beginnersbook.com/wp-content/uploads/2018/10/array.jpg)

## Time and Space complexity (Big-O)

-   Access an element: $O(1)$ time complexity and $O(1)$ space complexity
-   Insertion/Deletion:
    -   Start/Middle of an array: $O(n)$ both time and space complexity because the array has to shift to make room for new elements/fill the gap.
    -   End of an array: $O(1)$
-   Search:
    -   Linear search: $O(n)$ time complexity and $O(1)$ space complexity
    -   Binary search: $O(log(n))$ time complexity and $O(1)$ space complexity
-   Sort: refer to Big-O Cheatsheet
