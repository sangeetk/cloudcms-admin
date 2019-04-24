{{ template "layout/layout.tpl" . }}

{{ define "contents" }}

  {{ template "partial/flash.tpl" .}}
  
  <div class="columns is-multiline">
    <div class="column is-11">

<form method="POST" action="/admin/content/{{.Name}}" enctype="multipart/form-data">


      <div class="field is-horizontal">
        <div class="field-label is-normal">
          <label class="label">Header</label>
        </div>
        <div class="field-body">
          <div class="field is-expanded has-addons">
            <p class="control">
              <a class="button is-info"> ID </a>
            </p>
            <p class="control is-expanded">
              <input class="input" type="text" value="{{with .Content}}{{.id}}{{end}}" disabled>
              <input type="hidden" name="id" id="id" value="{{with .Content}}{{.id}}{{end}}">
            </p>
          </div>
          <div class="field is-expanded has-addons">
            <p class="control">
              <a class="button is-info"> Lang </a>
            </p>
            <p class="control is-expanded">
              {{ if .Content }}
              <input class="input" type="text" value="{{langCodeToName .Content.language}}" disabled>
              <input type="hidden" name="language" id="language" value="{{.Content.language}}">
              {{ else }}
              <input class="input" type="text" value="{{langCodeToName .LanguageCode}}" disabled>
              <input type="hidden" name="language" id="language" value="{{.LanguageCode}}">
              {{ end }}
            </p>
          </div>
          <div class="field is-expanded has-addons">
            <p class="control">
              <a class="button is-info"> Status </a>
            </p>
            <p class="control is-expanded">
              <input class="input" type="text" name="status" id="status" value="{{with .Content}}{{if .status}}{{status .status}}{{end}}{{end}}" disabled>
            </p>
          </div>
        </div>
      </div>
      
      <div class="field is-horizontal">
        <div class="field-label is-normal">
          <label class="label">Timestamps</label>
        </div>
        <div class="field-body">
          <div class="field is-expanded has-addons">
            <p class="control">
              <a class="button is-info"> C </a>
            </p>
            <p class="control is-expanded">
              <input class="input" type="text" value="{{with .Content}}{{with .created_at}}{{unixTimeToString .}}{{end}}{{end}}" disabled>
            </p>
          </div>
          <div class="field is-expanded has-addons">
            <p class="control">
              <a class="button is-info"> U </a>
            </p>
            <p class="control is-expanded">
              <input class="input" type="text" value="{{with .Content}}{{with .updated_at}}{{unixTimeToString .}}{{end}}{{end}}" disabled>
            </p>
          </div>
          <div class="field is-expanded has-addons">
            <p class="control">
              <a class="button is-info"> D </a>
            </p>
            <p class="control is-expanded">
              <input class="input" type="text" value="{{with .Content}}{{with .deleted_at}}{{unixTimeToString .}}{{end}}{{end}}" disabled>
            </p>
          </div>
        </div>
      </div>


      <div class="field is-horizontal">
        <div class="field-label is-normal">
          <label class="label">Slug</label>
        </div>
        <div class="field-body">
          <div class="field">
            <p class="control is-expanded">

              {{ if .TranslationSlug }}
              <input class="input" type="text" value="{{.TranslationSlug}}" disabled>
              <input type="hidden" name="translationslug" id="translationslug" value="{{.TranslationSlug}}">
              {{ else }}
              <input class="input" type="text" value="{{with .Content}}{{.slug}}{{end}}" disabled>
              <input type="hidden" name="slug" id="slug" value="{{with .Content}}{{.slug}}{{end}}">
              {{ end }}

            </p>
          </div>
        </div>
      </div>

      <br>

      {{ $currentCMS := .CurrentCMS }}
      {{ $fields := "" }}
      {{ $useasslug := "" }}
      {{ $useforslug := "" }}
      {{ $content := .Content }}
      {{ $newContent := false }}
      {{ if hasPrefix .Title "Add" }} {{ $newContent = true}} {{ end }}

      {{ range $f := .Fields }}
      {{ $name := lowercase $f.Name }}

      {{ $fields = appendField $fields $f }}
      {{ if $f.UseAsSlug }}
        {{ $useasslug = lowercase $f.Name }}
      {{ end }}
      {{ if $f.UseForSlug }}
        {{ $useforslug = appendField $useforslug $f }}
      {{ end }}

      {{ if eq $f.Widget "input" }}
      {{ if not $f.SkipHeader }}
      <div class="field is-horizontal">
        <div class="field-label is-normal">
          <label class="label">{{if $f.Heading}}{{$f.Heading}}{{else}}{{$f.Name}}{{end}}</label>
        </div>
        <div class="field-body">
        {{ end }}
          <div class="field has-addons">
            {{ if $f.HasLabel }}
            <p class="control">
              <a class="button is-disabled"> {{ $f.Name }} </a>
            </p>
            {{ end }}
            <div class="control is-expanded">
              <input class="input" type="text" id="{{$name}}" name="{{$name}}" 
                value="{{contentTextValue $content $name}}" placeholder="{{ $f.Helptext }}">
            </div>
          </div>
        {{ if not $f.SkipFooter }}
        </div>
      </div>
      {{ end }}
      {{ else if eq $f.Widget "textarea" }}
      <div class="field is-horizontal">
        <div class="field-label is-normal">
          <label class="label">{{if $f.Heading}}{{$f.Heading}}{{else}}{{$f.Name}}{{end}}</label>
        </div>
        <div class="field-body">
          <div class="field">
            <div class="control">
              <textarea class="textarea" id="{{$name}}" name="{{$name}}" value="{{contentTextValue $content $name}}"
                placeholder="{{ $f.Helptext }}"></textarea>
            </div>
          </div>
        </div>
      </div>
      {{ else if eq $f.Widget "richtext" }}
      <div class="field is-horizontal">
        <div class="field-label is-normal">
          <label class="label">{{if $f.Heading}}{{$f.Heading}}{{else}}{{$f.Name}}{{end}}</label>
        </div>
        <div class="field-body">
          <div class="field">
            <div class="control">
              <input type="hidden" id="{{$name}}" name="{{$name}}" value="{{contentTextValue $content $name}}">
              <trix-editor input="{{$name}}"></trix-editor>
            </div>
          </div>
        </div>
      </div>
      {{ else if eq $f.Widget "date" }}
      {{ if not $f.SkipHeader }}
      <div class="field is-horizontal">
        <div class="field-label is-normal">
          <label class="label">{{if $f.Heading}}{{$f.Heading}}{{else}}{{$f.Name}}{{end}}</label>
        </div>
        <div class="field-body">
        {{ end }}
          <div class="field has-addons">
            {{ if $f.HasLabel }}
            <p class="control">
              <a class="button is-disabled"> {{ $f.Name }} </a>
            </p>
            {{ end }}
            <div class="control is-expanded">
              <input class="input" type="text" name="{{$name}}" id="{{$name}}" 
              {{ if $newContent }} 
                value="{{ currentDate }}"
              {{ else }}
                value="{{ contentDateValue $content $name }}" 
              {{ end }}
              placeholder="{{$f.Helptext}}">
            </div>
          </div>
        {{ if not $f.SkipFooter }}
        </div>
      </div>
      {{ end }}
      {{ else if eq $f.Widget "file" }}
      <div class="field is-horizontal">
        <div class="field-label is-normal">
          <label class="label">{{ trimPrefix $f.Name "file:" }}</label>
        </div>
        <div class="field-body">
          <div class="field">
            <div class="control">
              <div class="file has-name">
                <label class="file-label">
                  {{ $file := contentFile $content $name }}
                  <input class="file-input" type="file" name="{{$name}}" id="{{$name}}">
                  <input type="hidden" name="{{$name}}.name" value="{{with $file}}{{.Name}}{{end}}">
                  <input type="hidden" name="{{$name}}.size" value="{{with $file}}{{.Size}}{{end}}">
                  <input type="hidden" name="{{$name}}.uri" value="{{with $file}}{{.URI}}{{end}}">
                  <input type="hidden" name="{{$name}}.type" value="{{$f.FileType}}">
                  <span class="file-cta">
                    <span class="file-icon"> <i class="fas fa-upload"></i> </span>
                    <span class="file-label"> Select {{ trimPrefix $f.Name "file:" }}... </span>
                  </span>
                  <span class="button is-white"> {{with $file}}{{.Name}}{{end}} </span>
                </label>
              </div>
            </div>

            <div class="control is-expanded">
              <br>
              {{with $file}}
                {{if eq .Type "image"}}
                  {{with .URI}}
                  <img src="{{ drive .}}">
                  {{end}}
                {{end}}
              {{end}}
            </div>

          </div>
        </div>
      </div>
      {{ else if eq $f.Widget "select" }}
      <div class="field is-horizontal">
        <div class="field-label is-normal">
          <label class="label">{{if $f.Heading}}{{$f.Heading}}{{else}}{{$f.Name}}{{end}}</label>
        </div>
        <div class="field-body">
          <div class="field has-addons">
            {{ if $f.Label }}
            <p class="control">
              <a class="button is-info"> {{ $f.Label }} </a>
            </p>
            {{ end }}
            <div class="control is-expanded">
              <div class="select is-normal">
                <select>
                  <option> -- None -- </option>
                  <option>Category 1</option>
                </select>
              </div>
            </div>
          </div>
        </div>
      </div>
      {{ else if eq $f.Widget "selectmultiple" }}
      <div class="field is-horizontal">
        <div class="field-label is-normal">
          <label class="label">{{if $f.Heading}}{{$f.Heading}}{{else}}{{$f.Name}}{{end}}</label>
        </div>
        <div class="field-body">
          <div class="field has-addons">
            {{ if $f.HasLabel }}
            <p class="control">
              <a class="button is-info"> {{ $f.Name }} </a>
            </p>
            {{ end }}
            <div class="control is-expanded">
              <div class="select is-normal is-multiple">
                <select multiple size="5">
                  <option> -- None -- </option>
                  <option>Category 1</option>
                  <option>Category 2</option>
                  <option>Category 3</option>
                  <option>Category 4</option>
                </select>
              </div>
            </div>
          </div>
        </div>
      </div>
      {{ else if eq $f.Widget "tags" }}
      {{ if not $f.SkipHeader }}
      <div class="field is-horizontal">
        <div class="field-label is-normal">
          <label class="label">{{if $f.Heading}}{{$f.Heading}}{{else}}{{$f.Name}}{{end}}</label>
        </div>
        <div class="field-body">
        {{ end }}
          <div class="field has-addons">
            {{ if $f.HasLabel }}
            <p class="control">
              <a class="button is-disabled"> {{ $f.Name }} </a>
            </p>
            {{ end }}
            <p class="control is-expanded">
              <input class="input" type="text" name="{{$name}}" id="{{$name}}" value="{{contentTagsValue $content $name}}" placeholder="{{$f.Helptext}}">
            </p>
          </div>
        {{ if not $f.SkipFooter }}
        </div>
      </div>
      {{ end }}
      {{ else if eq $f.Widget "checkbox" }}
      <div class="field is-horizontal">
        <div class="field-label is-normal">
          <label class="label">{{if $f.Heading}}{{$f.Heading}}{{else}}{{$f.Name}}{{end}}</label>
        </div>
        <div class="field-body">
          <div class="field has-addons">
            {{ if $f.HasLabel }}
            <p class="control">
              <a class="button is-info"> {{ $f.Name }} </a>
            </p>
            {{ end }}
            <div class="control is-expanded">
              <label class="checkbox">
                <input type="checkbox" checked="yes">
                {{ $f.Helptext }}
              </label>
            </div>
          </div>
        </div>
      </div>
      {{ else if eq $f.Widget "radio" }}
      <div class="field is-horizontal">
        <div class="field-label is-normal">
          <label class="label">{{if $f.Heading}}{{$f.Heading}}{{else}}{{$f.Name}}{{end}}</label>
        </div>
        <div class="field-body">
          <div class="field has-addons">
            {{ if $f.HasLabel }}
            <p class="control">
              <a class="button is-info"> {{ $f.Name }} </a>
            </p>
            {{ end }}
            <div class="control is-expanded">
              <label class="radio">
                <input type="radio" name="foobar">
                Foo
              </label>
              <label class="radio">
                <input type="radio" name="foobar" checked>
                Bar
              </label>
            </div>
          </div>
        </div>
      </div>
      {{ end }}

      {{ end }}
      
      <input type="hidden" name="fields" value="{{$fields}}">
      <input type="hidden" name="useasslug" value="{{$useasslug}}">
      <input type="hidden" name="useforslug" value="{{$useforslug}}">
      <br>

      <div class="field is-horizontal">
        <div class="field-label">
          <!-- Left empty for spacing -->
        </div>
        <div class="field-body">
          <div class="field">
            <div class="control">
              <button class="button is-primary">
                {{ .SubmitButton }}
              </button>
            </div>
          </div>
          {{ if hasPrefix .Title "Edit" }}
          <div class="field">
            <div class="control">
              <a class="button is-danger" href="/admin/content/{{.Name}}/delete?slug={{.Content.slug}}">
                Delete {{ title .Name }}
              </a>
            </div>
          </div>
          {{ end }}
        </div>
      </div>
      
</form>

    </div>
  </div>

<br>

{{ end }}