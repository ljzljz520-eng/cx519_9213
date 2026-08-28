package flow061
import("path/filepath";"testing";"soundspace/internal/model";"soundspace/internal/store")
func TestWorkflowSearchUpdatePublish(t *testing.T){s,_:=store.Open(filepath.Join(t.TempDir(),"s.db"));defer s.Close();f:=New(s);f.Register(model.Record{ID:"s",Title:"Search",Location:"L",Status:"draft"},"u");f.Review("s","u");f.Approve("s","u","ok");if e:=f.Update("s","Changed","note","u");e!=nil{t.Fatal(e)};if e:=f.Publish("s","u");e!=nil{t.Fatal(e)}}
func TestWorkflowImportReport(t *testing.T){s,_:=store.Open(filepath.Join(t.TempDir(),"i.db"));defer s.Close();f:=New(s);if e:=f.RegisterBatch([]model.Record{{ID:"i",Title:"Imported",Location:"L",Status:"draft"}},"import");e!=nil{t.Fatal(e)};r,_:=s.GetRecord("i");if r.Title!="Imported"{t.Fatal(r)}}
