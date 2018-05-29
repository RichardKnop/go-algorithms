# Quicksort

Worst case: O(n2)
Best case: O(n log n)
Average case: O(n log n)

Similar to selection sort but pivot is not always chosen as leftmost item which leads to worst case complexity for already sorted arrays.

Typically, quicksort is significantly faster in practice than other Θ(nlogn) algorithms, because its inner loop can be efficiently implemented on most architectures, and in most real-world data, it is possible to make design choices which minimize the probability of requiring quadratic time.

Quick sort is in place so low memory usage.