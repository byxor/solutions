(fn replicate [collection times]
  (if (not (= collection []))
    (concat
      (repeat times (first collection))
      (replicate (rest collection) times))))
