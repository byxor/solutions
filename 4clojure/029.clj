(fn get_the_caps [message]
  (apply str (re-seq #"[A-Z]" message)))
