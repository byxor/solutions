(fn reverse_ [collection]
  (cond
    (= (count collection) 0) (list)
    (= (count collection) 1) (list (first collection))
    :else (concat
        (reverse_ (rest collection))
        (list (first collection)))))
