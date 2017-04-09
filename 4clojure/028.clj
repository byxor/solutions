(fn flatten_ [collection]
  (if (= collection [])
    ()
    (let [head (first collection), tail (rest collection)]
      (if (sequential? head)
        (concat (flatten_ head) (flatten_ tail))
        (concat (list head) (flatten_ tail))))))
