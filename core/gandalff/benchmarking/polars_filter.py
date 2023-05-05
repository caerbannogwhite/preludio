print("# filter-polars.py", flush=True)

import os
import gc
import timeit
import polars as pl

from statistics import mean
from helpers import write_log, memory_usage, make_chk

ver = pl.__version__
git = ""
task = "filter"
solution = "polars"
fun = ".filter"
cache = "TRUE"
on_disk = "FALSE"
data_names = ["G1_1e5_1e2_0_0", "G1_1e6_1e2_0_0", "G1_1e7_1e2_0_0"]


for data_name in data_names:
    filepath = os.path.join("..", "testdata", data_name+".csv")
    print("loading dataset %s" % data_name, flush=True)

    with pl.StringCache():
        x = pl.read_csv(filepath, dtypes={"id4":pl.Int32, "id5":pl.Int32, "id6":pl.Int32, "v1":pl.Int32, "v2":pl.Int32, "v3":pl.Float64}, low_memory=True)
        # x["id1"] = x["id1"].cast(pl.Categorical)
        x["id1"].shrink_to_fit(in_place=True)
        # x["id2"] = x["id2"].cast(pl.Categorical)
        x["id2"].shrink_to_fit(in_place=True)
        # x["id3"] = x["id3"].cast(pl.Categorical)
        x["id3"].shrink_to_fit(in_place=True)

    in_rows = x.shape[0]
    # x = x.lazy()


    ###################     QUESTION 1   ###################

    question = "filtering: id6 > 500 or id1 == 'id024'"
    gc.collect()
    times = timeit.repeat(lambda: x.filter((x["id6"] > 500) | (x["id1"] == "id024")), repeat=5, number=10)
    t = mean(times)

    gc.collect()
    ans =  x.filter((x["id6"] > 500) | (x["id1"] == "id024"))
    print(ans.shape, flush=True)
    m = memory_usage()
    chk = [ans["v1"].cast(pl.Int64).sum()]

    write_log(task=task, data=data_name, in_rows=in_rows, question=question, out_rows=ans.shape[0], out_cols=ans.shape[1], solution=solution, version=ver, git=git, fun=fun, run=1, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=-1, on_disk=on_disk)
    del ans


    ###################     QUESTION 2   ###################

    question = "filtering: q2"
    gc.collect()
    times = timeit.repeat(lambda: x.filter((x["id6"] > 500) & (x["v3"] < 50) & ((x["id1"] == "id024") | (x["id2"] == "id024")) & (x["v1"] == 5) & (x["v2"] == 1)), repeat=5, number=10)
    t = mean(times)

    gc.collect()
    ans =  x.filter((x["id6"] > 500) & (x["v3"] < 50) & ((x["id1"] == "id024") | (x["id2"] == "id024")) & (x["v1"] == 5) & (x["v2"] == 1))
    print(ans.shape, flush=True)
    m = memory_usage()
    chk = [ans["v1"].cast(pl.Int64).sum()]

    write_log(task=task, data=data_name, in_rows=in_rows, question=question, out_rows=ans.shape[0], out_cols=ans.shape[1], solution=solution, version=ver, git=git, fun=fun, run=1, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=-1, on_disk=on_disk)
    del ans

