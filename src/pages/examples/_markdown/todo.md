### Folder hierarchy

```bash
app
 ├─ index.html
 └─ todo
    ├─ render.go
    ├─ handle.go
    └─ todo.html
```

### render.go

```go
package todo

import "github.com/caleb-sideras/gox/.gox/render"

func TodoSsr_() render.StaticF {
	return render.StaticF{[]string{"app/todo/ssr.html"}, nil, "todo-ssr"}
}
```

### handle.go

```go
package todo

import (
	"html/template"
	"net/http"
	"time"
)

func HandleAddTask_(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	task := r.PostFormValue("task")
	tmpl := template.Must(template.ParseFiles("pages/todo/ssr.html"))
	tmpl.ExecuteTemplate(w, "task-list-element", Task{Text: task})
}
```

### ssr.html

```html
{{`{{ block "todo-ssr" .}}`}}
<form hx-post="/todo/handle-add-task" hx-target="#task-list" hx-swap="beforeend" hx-indicator="#spinner"
  class="flex justify-between items-center w-full gap-4">
  <md-outlined-text-field label="Add a new task" type="text" name="task" id="task" class="form-control">
    <md-circular-progress slot="trailingicon" indeterminate class="htmx-indicator" id="spinner" role="status"
      aria-hidden="true">
    </md-circular-progress>
  </md-outlined-text-field>

  <md-filled-button type="submit" class="items-center h-full">
    <md-icon>
      <span class="material-symbols-outlined"
        style="font-variation-settings: 'FILL' 1, 'wght' 400, 'GRAD' 0, 'opsz' 48;">
        add
      </span>
    </md-icon>
  </md-filled-button>
</form>

<ul class="list-group overflow-auto flex-1 bg-secondary-container text-on-secondary-container rounded-md"
  id="task-list">
  {{`{{ range .Tasks}}`}}
  {{`{{ block "task-list-element" .}}`}}
  <li class="w-full justify-between align-center flex p-4">
    <span class='list-group-item my-auto'>{{`{{.Text}}`}}</span>
    <md-checkbox></md-checkbox>
  </li>
  <md-divider inset></md-divider>
  {{`{{ end }}`}}
  {{`{{ end }}`}}
</ul>
{{`{{end}}`}}
```