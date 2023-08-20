print("# filter-pandas.py", flush=True)

import os
import gc
import timeit
import pandas as pd

from statistics import mean
from helpers import write_log, memory_usage, make_chk

# other questions ans info here
# https://github.com/h2oai/db-benchmark/tree/master

ver = pd.__version__
git = ""
task = "groupby"
solution = "pandas"
fun = ".groupby"
cache = "TRUE"
on_disk = "FALSE"
data_names = [
    "G1_1e4_1e2_0_0", "G1_1e5_1e2_0_0", "G1_1e6_1e2_0_0", "G1_1e7_1e2_0_0",
    "G1_1e4_1e2_10_0", "G1_1e5_1e2_10_0", "G1_1e6_1e2_10_0", "G1_1e7_1e2_10_0"
]

for data_name in data_names:
    filepath = os.path.join("..", "testdata", data_name+".csv")
    print("loading dataset %s" % data_name, flush=True)

    x = pd.read_csv(filepath)


    ###################     QUESTION 1   ###################

    question = "sum v1 by id1"
    gc.collect()
    times = timeit.repeat(lambda: x.groupby('id1', as_index=False, sort=False, observed=True, dropna=False).agg({'v1':'sum'}), repeat=5, number=1)
    t = mean(times)

    gc.collect()
    ans = x.groupby('id1', as_index=False, sort=False, observed=True, dropna=False).agg({'v1':'sum'})
    print(ans.shape, flush=True)
    m = memory_usage()
    chk = [ans['v1'].sum()]

    write_log(task=task, data=data_name, in_rows=x.shape[0], question=question, out_rows=ans.shape[0], out_cols=ans.shape[1], solution=solution, version=ver, git=git, fun=fun, run=1, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=-1, on_disk=on_disk)
    del ans


    ###################     QUESTION 2   ###################

    question = "sum v1 by id1:id2"
    gc.collect()
    times = timeit.repeat(lambda: x.groupby(['id1','id2'], as_index=False, sort=False, observed=True, dropna=False).agg({'v1':'sum'}), repeat=5, number=1)
    t = mean(times)

    gc.collect()
    ans = x.groupby(['id1','id2'], as_index=False, sort=False, observed=True, dropna=False).agg({'v1':'sum'})
    print(ans.shape, flush=True)
    m = memory_usage()
    chk = [ans['v1'].sum()]

    write_log(task=task, data=data_name, in_rows=x.shape[0], question=question, out_rows=ans.shape[0], out_cols=ans.shape[1], solution=solution, version=ver, git=git, fun=fun, run=1, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=-1, on_disk=on_disk)
    del ans


    ###################     QUESTION 3   ###################

    question = "sum v1 mean v3 by id3"
    gc.collect()
    times = timeit.repeat(lambda: x.groupby('id3', as_index=False, sort=False, observed=True, dropna=False).agg({'v1':'sum', 'v3':'mean'}), repeat=5, number=1)
    t = mean(times)

    gc.collect()
    ans = x.groupby('id3', as_index=False, sort=False, observed=True, dropna=False).agg({'v1':'sum', 'v3':'mean'})
    print(ans.shape, flush=True)
    m = memory_usage()
    chk = [ans['v1'].sum(), ans['v3'].sum()]

    write_log(task=task, data=data_name, in_rows=x.shape[0], question=question, out_rows=ans.shape[0], out_cols=ans.shape[1], solution=solution, version=ver, git=git, fun=fun, run=1, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=-1, on_disk=on_disk)
    del ans


    ###################     QUESTION 4   ###################

    question = "mean v1:v3 by id4"
    gc.collect()
    times = timeit.repeat(lambda: x.groupby('id4', as_index=False, sort=False, observed=True, dropna=False).agg({'v1':'mean', 'v2':'mean', 'v3':'mean'}), repeat=5, number=1)
    t = mean(times)

    gc.collect()
    ans = x.groupby('id4', as_index=False, sort=False, observed=True, dropna=False).agg({'v1':'mean', 'v2':'mean', 'v3':'mean'})
    print(ans.shape, flush=True)
    m = memory_usage()
    chk = [ans['v1'].sum(), ans['v2'].sum(), ans['v3'].sum()]

    write_log(task=task, data=data_name, in_rows=x.shape[0], question=question, out_rows=ans.shape[0], out_cols=ans.shape[1], solution=solution, version=ver, git=git, fun=fun, run=1, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=-1, on_disk=on_disk)
    del ans


    ###################     QUESTION 5   ###################

    question = "sum v1:v3 by id6"
    gc.collect()
    times = timeit.repeat(lambda: x.groupby('id6', as_index=False, sort=False, observed=True, dropna=False).agg({'v1':'sum', 'v2':'sum', 'v3':'sum'}), repeat=5, number=1)
    t = mean(times)

    gc.collect()
    ans = x.groupby('id6', as_index=False, sort=False, observed=True, dropna=False).agg({'v1':'sum', 'v2':'sum', 'v3':'sum'})
    print(ans.shape, flush=True)
    m = memory_usage()
    chk = [ans['v1'].sum(), ans['v2'].sum(), ans['v3'].sum()]

    write_log(task=task, data=data_name, in_rows=x.shape[0], question=question, out_rows=ans.shape[0], out_cols=ans.shape[1], solution=solution, version=ver, git=git, fun=fun, run=1, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=-1, on_disk=on_disk)
    del ans