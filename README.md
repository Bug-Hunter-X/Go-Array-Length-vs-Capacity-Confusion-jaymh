# Go Array Length vs Capacity Confusion

This example demonstrates a potential source of confusion for Go programmers: the relationship between the length and capacity of an array.

In Go, arrays have a fixed size.  The length and capacity are always identical. Attempting to modify the length of an array after creation will result in errors. This is in contrast to slices which can dynamically grow and shrink.

The `len()` and `cap()` functions return the same value for arrays.

**Bug:** The code may lead to a misunderstanding of how arrays are handled in Go.

**Solution:** The solution clarifies the distinction between arrays and slices and demonstrates how to use them appropriately.