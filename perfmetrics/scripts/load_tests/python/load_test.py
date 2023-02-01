from absl import logging
import importlib
import inspect
import logging
import os
import sys
import time
import argparse
from load_generator import load_generator as lg
from load_generator import task


class LoadGeneratorForReadAndWriteTask(lg.LoadGenerator):

  TASK_TYPES = ['read', 'write']

  def pre_load_test(self, **kwargs):
    # only run custom logic for read and write tasks
    if getattr(kwargs['task'], 'TASK_TYPE', '').lower() not in self.TASK_TYPES:
      return

    if not hasattr(kwargs['task'], 'FILE_PATH_FORMAT') or \
        not hasattr(kwargs['task'], 'FILE_SIZE'):
      raise Exception("Task of types - read or write must have FILE_PATH_FORMAT "
                      "and FILE_SIZE (in bytes) attributes set.")

    # Create one file per process for read and write tasks.
    file_path_format = getattr(kwargs['task'], 'FILE_PATH_FORMAT')
    file_size = getattr(kwargs['task'], 'FILE_SIZE')
    logging.info("One file is created per process of size {0} using the format "
                 "{1}".format(file_size, file_path_format))
    for process_num in range(self.num_processes):
      file_path = file_path_format.format()
      if os.path.exists(file_path) and os.path.getsize(file_path) == file_size:
        continue
      logging.info("Creating file {0} of size {1}.".format(file_path, file_size))
      with open(file_path, "wb") as fp:
        fp.truncate(file_size)

    return


  def post_load_test(self, observations, output_dir="./", dump_metrics=True,
      print_metrics=True, **kwargs):

    metrics = super().post_load_test(observations, output_dir, dump_metrics,
                                     print_metrics)
    # only run custom logic for read and write tasks
    if getattr(kwargs['task'], 'TASK_TYPE', '').lower() not in self.TASK_TYPES:
      return metrics

    metrics = metrics['metrics']

    # compute bandwidth from task results
    total_io_bytes = sum([task_result[4]
                          for task_result in observations['tasks_results_queue']])
    avg_computed_net_bw = total_io_bytes / metrics['actual_run_time']
    avg_computed_net_bw = avg_computed_net_bw * (1024 * 1024)
    metrics.update({'avg_computed_net_bw': avg_computed_net_bw})

    # dump metrics
    self._dump_results_into_json(metrics, output_dir)

    # print additional metrics
    print("\nNetwork bandwidth (computed by Sum(task response) / actual run time.")
    print("\n\tAvg. bandwidth (MiB/sec): ", metrics['avg_computed_net_bw'])
    return {'metrics': metrics}


def parse_args():
  # ToDo: Correct the description and defaults.
  parser = argparse.ArgumentParser(description='Load testing using multiprocessing')
  parser.add_argument('--task-file-path', type=str,
                      help='Path to task file.')
  parser.add_argument('--task-names', default=[],
                      action=lambda nms: nms.replace(" ", "").split(","),
                      help="")
  parser.add_argument('--output-dir', type=str, default='./output/',
                      help='Path to task file.')
  parser.add_argument('--num-processes', type=int, default=1,
                      help='Number of processes to use.')
  parser.add_argument('--num-threads', type=int, default=1,
                      help='Number of threads to use in each process.')
  parser.add_argument('--run-time', type=int, default=60,
                      help='Number of seconds to run the load test.')
  parser.add_argument('--num-tasks', type=int, default=sys.maxsize,
                      help='Number of seconds to run the load test.')
  parser.add_argument('--num-tasks-per-thread', type=int, default=sys.maxsize,
                      help='')
  parser.add_argument('--observation-interval', type=int,
                      default=lg.MIN_OBSERVATION_INTERVAL_IN_SECS,
                      help='')
  parser.add_argument('--cooling-time', type=int, default=10,
                      help='')
  parser.add_argument('--only-print',  action='store_true',
                      help='')
  args = parser.parse_args()
  return args


def main():
  args = parse_args()

  logging.info("Initialising Load Generator...")
  lg_obj = LoadGeneratorForReadAndWriteTask(num_processes=args.num_processes,
                               num_threads=args.num_threads,
                               run_time=args.run_time,
                               num_tasks_per_thread=args.num_tasks_per_thread,
                               num_tasks=args.num_tasks,
                               observation_interval=args.observation_interval)

  logging.info("Starting load generation...")
  args.task_file_path = args.task_file_path.replace(".py", "")
  for name, cls in inspect.getmembers(importlib.import_module(args.task_file_path), inspect.isclass):

    # Skip classes imported in the task file
    if cls.__module__ != args.task_file_path:
      continue
    # Skip classes that are not of type (Todo: add reference here)
    if type(cls) != type(task.LoadTestTask):
      continue
    # Skip if user only wants to run for a specific task
    if len(args.task_names) and cls.TASK_NAME not in args.task_names:
      continue

    task_obj = cls()
    logging.info("Running pre load test task for: ", cls.TASK_NAME)


    logging.info("Generating load for: ", cls.TASK_NAME)

    observations = lg_obj.generate_load(task_obj)

    output_dir = os.path.join(args.output_dir, cls.TASK_NAME)
    if not os.path.exists(output_dir):
      os.makedirs(output_dir)

    logging.info("Running post load test task for: ", cls.TASK_NAME)
    dump_metrics = not args.only_print
    metrics = lg_obj.post_load_test(observations, output_dir=output_dir,
                                    dump_metrics=dump_metrics,
                                    print_metrics=True)

    logging.info("Load test completed for task: ", cls.TASK_NAME)

    logging.info("Sleeping for {0} seconds...".format(args.cooling_time))
    time.sleep(args.cooing_time)

  return

if __name__ == "__main__":
  main()