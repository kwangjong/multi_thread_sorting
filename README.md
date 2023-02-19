# Multi-thread Sorting

This is a parallel implementation of the Quick Sort algorithm in Go, which takes advantage of multiple threads to sort large arrays more efficiently.

## Usage
To run the program, use the following command:
```
go run main.go -d <size>
```
where <size> is the number of int32 values to sort (default is 10,000,000).

## Output
The program will output the results of two different sorts: a single-threaded Quick Sort and a multi-threaded Quick Sort. The results include whether the array is sorted and the duration of the sort.

## Implementation
The program creates a random array of int32 values of the specified size, and then runs two different implementations of the Quick Sort algorithm: a single-threaded implementation and a multi-threaded implementation.

The multi-threaded implementation uses a sync.WaitGroup to keep track of the number of threads that are actively sorting, and limits the number of threads to the number of CPU cores available on the machine.

The program also includes an insertionSort() function to sort sub-arrays of size 15 or less, as this is generally faster than using the Quick Sort algorithm on small sub-arrays.

## Performance
The multi-threaded implementation generally performs faster than the single-threaded implementation on large arrays, as it takes advantage of multiple CPU cores to sort the data more quickly. On my machine, the single-threaded Quick Sort took `1.325355604 seconds` and the multi-threaded Quick Sort took `214.76232 milliseconds`. However, the actual speedup will depend on the specific hardware being used.

Note that the implementation in this program is not necessarily the fastest possible implementation of parallel Quick Sort in Go, and there are likely many ways to further optimize the algorithm for specific use cases.
