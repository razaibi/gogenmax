## gogenmax

A Golang based implementation of genmax.

#### Commands

| Command | Description |
|:---|:---|
|gmx init | Create a new project.|
|gmx run \<workflow-name> | Run a workflow and generate your files.|

#### Extensions supported in templates

The following methods are supported in the template:

| Description | Usage |
|:---|:---|
|Pluralize| `{{ "dog" \| pluralize }}`|
|Kebab Case| `{{ "Hello World" \| kebabcase }}`|
|Camel case| `{{ "Hello World" \| camelcase }}`|
|Snake case| `{{ "Hello World" \| snakecase }}`|
|Pascale case| `{{ "hello world" \| pascalecase }}`|
|UUID Generation| `{{ "" \| uuid }}`|
