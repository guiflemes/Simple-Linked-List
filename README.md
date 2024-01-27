# Simple Linked List Implementation

## Overview

This repository contains a fundamental implementation of a simple linked list in Go. It provides essential functionality for creating, manipulating, and iterating over a singly linked list data structure.

## Key Features

* Simple Linked List Structure: Nodes with forward pointers only.
  * Basic Operations:
  * Create a new list: New(elements []int) *List
  * Add elements: Push(element int) to add at the end.
  * Remove elements: Pop() (int, error) to remove from the end.
  * Get list size: Size() int
  * Convert to array: Array() []int
  * Reverse the list: Reverse() *List

## Running Tests

```go test -v --bench . --benchmem```
