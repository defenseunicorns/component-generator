# component-generator
Generate and aggregate OSCAL component definition files


## Commands

### Aggregate
An automated workflow for aggregation of local and remote OSCAL component-definition files into a single concise file. Each file that is included is converted to a component of a greater component definition file. 

#### Where is this useful?
If you are the maintainer or creator of a platform or system that is comprised of multiple components. Ideally those components would have their own unique component-definition files for compliance related material that you could utilize to produce an aggregate inheritance of said material in a reproducible fashion. 

We can create a configuration file that outlines pertinent platform/system component-definition data in a declarative fashion. This can then perform both retrieval and ingestion of files for aggregation. All components may not have an existing OSCAL component-definition, so we would allow users to create and reference these local copies for aggregation as well. 

#### Platform/System lifecycle

- Create a yaml configuration file with required information
    - Component name
    - etc
- Populate a field of local/remote component-definition files for aggregation
- Execute `component-generator aggregate config.yaml`

#### Execution 
What does the configuration file look like?

```yaml
name: my-generated-file.yaml
## Top-level metadata field - OSCAL compliant schema
metadata:
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
components:
    local:
    - name: test/input/jaeger-component-definition.yaml

    remote:
    - git: https://repo1.dso.mil/big-bang/apps/core/kiali.git@1.60.0-bb.2
      path: oscal-component.yaml
    - git: https://repo1.dso.mil/big-bang/apps/core/monitoring.git@43.1.2-bb.1
      path: oscal-component.yaml
```

The command should always default to the name provided in the configuration - unless it is empty

#### Hashes
Add functionality after aggregation is stable to allow the hash identification of the file in the declaration.
IE:

```yaml
name: my-generated-file.yaml
components:
    local:
    - name: component-1.yaml
      hash: <known hash>
```

## TODO
- Add ability to override version on the Command Line
- Proper error checking everywhere
- consistent logging