# -*- coding: utf-8 -*-

import numpy as np
import pandas as pd

print("hello world")


s = pd.Series([1, 3, 5, np.nan, 6, 8])

dates = pd.date_range('20130101', periods=6)
dates
