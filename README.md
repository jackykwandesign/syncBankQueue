# Sync Queue on async operation

```
force multiple call to form a queue to access protected data, such as account balance, which allow only 1 program access it

1. Call Instance
2. Call Instance.Start(), run immediately if empty queue, else wait until queue clear
3. Call Instance.End() to release the queue lock
```