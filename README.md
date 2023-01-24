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
    - This would ingest the configuration, retrieve required children component files, and produce a platform/system component-definition.
    - If an existing `component-name.yaml` file existed, it would perform UUID comparison against all children
        - If nothing has changed, it would re-use the existing UUID for the parent
        - If there has been changes, it would generate a new UUID/applicable timestamps and replace the file

#### Execution 
What does the configuration file look like?

```yaml
name: my-generated-file.yaml
components:
    local:
    - name: component-1.yaml
    - name: component-3.yaml

    remote:
    - git: github.com/my-application.git@v0.0.1
      path: ./oscal/component-definition.yaml
    - git: github.com/my-other-application.git@v0.1.2
      path: component-definition.yaml
```

The command should always default to stdout unless --output-file is explicitly defined.

If no file with the filename provided with --output-file is defined, create it with a new UUID.
- Then we will make an assumption that this should always change when undefined.

If a file with the same name is identified - we will perform some basic validation to ensure it is a prior generated file.
- Then we will store the existing UUID for the parent and all children in memory to start - perferrably this is a map?
- Then we will begin the collection of artifacts - foreach component - retrieve and identify UUID
- With all UUID's - determine if there was an update (new UUID) - if even 1 changed, we need a new parent UUID
- perform the aggregation - [example](https://repo1.dso.mil/big-bang/bigbang/-/blob/master/oscal-component.yaml)
- generate new file


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

This would allow additional checking for files that were incorrectly modified without update to UUID.
This could result in either a warning being produced or a failure?