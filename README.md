# component-generator

`component-generator` aggregates local and remote OSCAL component-definition files into a single concise file. Each file that is included is converted to a component of a greater component definition file.

## Where is this useful?

If you are the maintainer or creator of a platform or system that is comprised of multiple components, ideally those components would have their own unique component-definition files for compliance related material that can be used to produce an aggregate inheritance of said material in a reproducible fashion.

We can create a configuration file that outlines pertinent platform/system component-definition data in a declarative fashion. This can then perform both retrieval and ingestion of files for aggregation. All components may not have an existing OSCAL component-definition, so we would allow users to create and reference these local copies for aggregation as well.

### Usage

#### Build the binary from source

From the root of the repository, use the Makefile to build the binary

```bash
make build
```

This outputs the binary to `./bin/component-generator`

#### Create a configuration file

A YAML file is used to define the metadata and components of your OSCAL component definition file

Create a file named `oscal-components.yaml` with the following contents:

```yaml
name: my-generated-file.yaml # Name of the generated file
metadata: # OSCAL-compliant metadata
  title: my-oscal-document
  version: 0.0.1
  oscal-version: 1.0.4
  parties:
    - uuid: FFA360E2-0566-46AB-8982-2CCB787B78E3
      name: My Organization
      links:
      - href: https://myorganization.com
        rel: website
      type: organization
components: # Define paths on the local filesystem or remote paths in git repositories to OSCAL component definition files
    local:
    - name: testdata/input/jaeger-component-definition.yaml

    remote:
    - git: https://repo1.dso.mil/big-bang/apps/core/kiali.git@1.60.0-bb.2
      path: oscal-component.yaml
    - git: https://repo1.dso.mil/big-bang/apps/core/monitoring.git@43.1.2-bb.1
      path: oscal-component.yaml
```

#### Generate an OSCAL component definition file

```bash
./bin/component-generator aggregate --input oscal-components.yaml
```

You can now view the generated `my-generated-file.yaml`
