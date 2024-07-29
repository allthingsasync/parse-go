package text

import("testing")

func TestSourceText(t *testing.T) {
    st := NewSource("++-*")
    
    if st.Length() != 4 {
        t.Error("Length mismatch")
    }

    if st.IsEnd() {
        t.Error("Unexpected IsEnd")
    }
    
    ch := st.Eat()

    for ch != -1 {
        t.Logf("%c", ch)
        ch = st.Eat()
    }

    if !st.IsEnd() {
        t.Error("IsEnd Expected")
    }
}
