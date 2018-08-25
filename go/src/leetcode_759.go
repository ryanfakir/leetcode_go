func employeeFreeTime(schedule [][]Interval) []Interval {
    var ins []Interval
    for _, v := range schedule {
        ins = append(ins, v...)
    }
    sort.Sort(Intervals(ins))
    var res []Interval
    t := ins[0]
    for _, v := range ins {
        if t.End < v.Start {
            res = append(res, Interval{t.End, v.Start})
            t = v
        } else {
            if t.End < v.End {
                t = v
            }
        }
    }
    return res
}

type Intervals []Interval

func (is Intervals) Less(i, j int) bool {return is[i].Start < is[j].Start}
func (is Intervals) Len() int {return len(is)}
func (is Intervals) Swap(i, j int) {is[i], is[j] = is[j], is[i]}
