# Concurrency vs. Parallelism

## Concurrency

- **Definition**: Concurrency refers to the ability of a system to manage multiple tasks at the same time. It allows portions of different tasks to be executed in overlapping time periods within a single thread.
  
- **How It Works**:
  - In a concurrent system, a single thread executes part of one task, pauses, then executes part of another task, and continues this pattern.
  - Tasks are not completed sequentially; instead, they share the same thread and switch between tasks.

- **Example**: Imagine a chef preparing two dishes simultaneously. The chef may chop vegetables for the first dish, then pause to stir the sauce for the second dish, and then return to chopping.

## Parallelism

- **Definition**: Parallelism is the simultaneous execution of multiple tasks by utilizing separate threads or processors, allowing tasks to be completed at the same time.

- **How It Works**:
  - In a parallel system, each task is assigned to a separate thread or processor.
  - Tasks are executed independently and concurrently, allowing for greater efficiency and speed.

- **Example**: Using the same chef analogy, if two chefs are working together in the kitchen—one chef is preparing dish X while the other chef is preparing dish Y simultaneously, each chef works independently to complete their respective dish.

## Key Differences
| Aspect            | Concurrency                            | Parallelism                         |
|-------------------|---------------------------------------|------------------------------------|
| Execution         | Multiple tasks progress together in overlapping time | Multiple tasks executed simultaneously in separate threads |
| Resource Utilization | Shares the same resources (e.g., CPU thread) | Utilizes multiple resources (e.g., multiple CPU cores) |
| Goal              | Improve responsiveness and resource utilization | Increase throughput and reduce execution time |
