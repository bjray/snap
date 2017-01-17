# Plugin Status

Not all plugins are designed with the same level of readiness in mind. To visualize that fact while also allowing everyone to share their work, we have introduced the concept of "tiers" of plugins. These tiers are communicated at statuses which will be visualized everywhere you see the [Plugin Catalog](PLUGIN_CATALOG.md).

### What this is

* A way to quickly see the highest quality plugins
* A way to encourage contribution to existing plugins instead of forking

### What this is not

* Not a definition of the plugin version - that's done through `releases`
* Not a contract of support - us maintainers of Snap do our best but make no promises

### Status Tiers

Plugins will fall into one of the four tiers. Here are their brief overviews:
* **Supported** - Designed by Intel with the intent of supporting customers
* **Approved** - Vetted by Snap maintainers to meet our best practices for design
* **Experimental** - Early plugins ready for testing but not known to work as intended
* **Unlabeled** - Shared for reference or extension

Note that plugins can and likely will be demoted if they fall significantly behind best practices of the project. I know that's vague, but it will have to do for this first version :grin:

## Plugin Status Matrix

All plugins meet a set of requirements to be included at its status level.

| Feature                        |     Unlabeled      |    Experimental    |      Approved      |     Supported      |
|:-------------------------------|:------------------:|:------------------:|:------------------:|:------------------:|
| Naming (no conflict)           | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| Includes README.md             |                    | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| Active Maintainers             |                    | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| Release includes binary        |                    |                    | :white_check_mark: | :white_check_mark: |
| Known to Compile               |                    |                    | :white_check_mark: | :white_check_mark: |
| Known to Load in Snap          |                    |                    | :white_check_mark: | :white_check_mark: |
| Uses Snap Plugin Library       |                    |                    | :white_check_mark: | :white_check_mark: |
| Reasonable Test Coverage       |                    |                    | :white_check_mark: | :white_check_mark: |
| Includes Example Tasks         |                    |                    | :white_check_mark: | :white_check_mark: |
| Includes Dependency Management |                    |                    | :white_check_mark: | :white_check_mark: |
| Includes CI Status             |                    |                    | :white_check_mark: | :white_check_mark: |
| Includes License               |                    |                    | :white_check_mark: | :white_check_mark: |
| Built and Supported by Intel   |                    |                    |                    | :white_check_mark: |

## Supported Plugins

These are our premier plugins for the Snap telemetry framework and its users. These are designed to follow all of our recommended practices. Issues are also closely monitored by maintainers. While not all plugins Intel contributes to Snap will meet our Supported status, only plugins designed by Intel contributors will be eligible to be Supported.

We prefer to not have other repositories that overlap with Supported plugins and suggest contributing to the existing version to help keep this list small and effective. Like all plugins, community contribution is welcome.


## Approved Plugins

These are primarily community-contributed plugins that meet or exceed the project's best practices. These plugins have been vetted by Snap maintainers as of the date listed in the Plugin Catalog. They are excellent references are quite likely ready for use in your own use of Snap.

We prefer to not have other repositories that overlap with Approved plugins and suggest contributing to the existing version to help keep this list small and effective. Don't be shy about reaching out to existing plugin authors to see if you can help improve upon it.

## Experimental

These are plugins being developed that are not yet complete, but are shared with the community for feedback and testing. This of this level as an incubation phase that should move toward Approved or drop down to Unlabeled.

We prefer to not have other repositories that overlap with Experimental plugins and suggest contributing to the existing version to help keep this list small and effective. Don't be shy about reaching out to existing plugin authors to see if you can help improve upon it.


## All Other Plugins (Unlabeled)

In early development, for reference and community supported. Does not necessarily follow best practices for plugin development. We welcome forking and working toward Approved status.
