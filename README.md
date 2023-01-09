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
        - If there has been changes, it would generate a new UUID and applicable timestamps