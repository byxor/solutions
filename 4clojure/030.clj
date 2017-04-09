; I spent a long time trying to clean this up.
; It's hard when extracting functions isn't allowed.

(fn un-dupe [something]
  (let [collection (seq something)
        _1st (first collection)
        _2nd (second collection)]
    (if (= (count (set collection)) 1)
      (list _1st)
      (if (= (count collection) 2)
        (if (= _1st _2nd) () collection)
        (if (= _1st _2nd)
          (un-dupe (rest collection))
          (conj (un-dupe (rest collection)) _1st))))))
