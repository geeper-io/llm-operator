# API Reference

Packages:
- [llm.geeper.io/v1alpha1](#llm-geeper-io-v1alpha1)

# llm.geeper.io/v1alpha1

Resource Types:
- [LMDeployment](#lmdeployment)
## LMDeployment

LMDeployment is the Schema for the lmdeployments API

| Name | Type | Description | Required |
|------|------|-------------|----------|
| apiVersion | string | llm.geeper.io/v1alpha1 | true |
| kind | string | LMDeployment | true |
| [metadata](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta) | object | Refer to the Kubernetes API docs for the `metadata` field. | true |
| [spec](#lmdeploymentspec) | object | LMDeploymentSpec defines the desired state of Deployment | false |
| [status](#lmdeploymentstatus) | object | LMDeploymentStatus defines the observed state of Deployment | false |
### LMDeployment.spec

LMDeploymentSpec defines the desired state of Deployment

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [ollama](#lmdeploymentspecollama) | object | Ollama defines the Ollama deployment configuration | false |
| [openwebui](#lmdeploymentspecopenwebui) | object | OpenWebUI defines the OpenWebUI deployment configuration | false |
| [tabby](#lmdeploymentspectabby) | object | Tabby defines the Tabby deployment configuration | false |
| [vllm](#lmdeploymentspecvllm) | object | VLLM defines the vLLM deployment configuration | false |
### LMDeployment.spec.ollama

Ollama defines the Ollama deployment configuration

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [affinity](#lmdeploymentspecollamaaffinity) | object | Affinity defines pod affinity and anti-affinity rules for Ollama pods | false |
| enabled | boolean | Enabled determines if vLLM should be deployed instead of Ollama | false |
| image | string | Image is the Ollama container image to use (including tag) | false |
| models | []string | Models is the list of models to deploy with Ollama | false |
| replicas | integer | Replicas is the number of Ollama pods to run<br/>*Format*: int32<br/> | false |
| [resources](#lmdeploymentspecollamaresources) | object | Resources defines the resource requirements for Ollama pods | false |
| [service](#lmdeploymentspecollamaservice) | object | Service defines the service configuration for Ollama | false |
### LMDeployment.spec.ollama.affinity

Affinity defines pod affinity and anti-affinity rules for Ollama pods

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [nodeAffinity](#lmdeploymentspecollamaaffinitynodeaffinity) | object | Describes node affinity scheduling rules for the pod. | false |
| [podAffinity](#lmdeploymentspecollamaaffinitypodaffinity) | object | Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)). | false |
| [podAntiAffinity](#lmdeploymentspecollamaaffinitypodantiaffinity) | object | Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)). | false |
### LMDeployment.spec.ollama.affinity.nodeAffinity

Describes node affinity scheduling rules for the pod.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [preferredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspecollamaaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindex) | []object | The scheduler will prefer to schedule pods to nodes that satisfy<br/>the affinity expressions specified by this field, but it may choose<br/>a node that violates one or more of the expressions. The node that is<br/>most preferred is the one with the greatest sum of weights, i.e.<br/>for each node that meets all of the scheduling requirements (resource<br/>request, requiredDuringScheduling affinity expressions, etc.),<br/>compute a sum by iterating through the elements of this field and adding<br/>"weight" to the sum if the node matches the corresponding matchExpressions; the<br/>node(s) with the highest sum are the most preferred. | false |
| [requiredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspecollamaaffinitynodeaffinityrequiredduringschedulingignoredduringexecution) | object | If the affinity requirements specified by this field are not met at<br/>scheduling time, the pod will not be scheduled onto the node.<br/>If the affinity requirements specified by this field cease to be met<br/>at some point during pod execution (e.g. due to an update), the system<br/>may or may not try to eventually evict the pod from its node. | false |
### LMDeployment.spec.ollama.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index]

An empty preferred scheduling term matches all objects with implicit weight 0<br/>(i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [preference](#lmdeploymentspecollamaaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindexpreference) | object | A node selector term, associated with the corresponding weight. | true |
| weight | integer | Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.<br/>*Format*: int32<br/> | true |
### LMDeployment.spec.ollama.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].preference

A node selector term, associated with the corresponding weight.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecollamaaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindexpreferencematchexpressionsindex) | []object | A list of node selector requirements by node's labels. | false |
| [matchFields](#lmdeploymentspecollamaaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindexpreferencematchfieldsindex) | []object | A list of node selector requirements by node's fields. | false |
### LMDeployment.spec.ollama.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].preference.matchExpressions[index]

A node selector requirement is a selector that contains values, a key, and an operator<br/>that relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The label key that the selector applies to. | true |
| operator | string | Represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt. | true |
| values | []string | An array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. If the operator is Gt or Lt, the values<br/>array must have a single element, which will be interpreted as an integer.<br/>This array is replaced during a strategic merge patch. | false |
### LMDeployment.spec.ollama.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].preference.matchFields[index]

A node selector requirement is a selector that contains values, a key, and an operator<br/>that relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The label key that the selector applies to. | true |
| operator | string | Represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt. | true |
| values | []string | An array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. If the operator is Gt or Lt, the values<br/>array must have a single element, which will be interpreted as an integer.<br/>This array is replaced during a strategic merge patch. | false |
### LMDeployment.spec.ollama.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution

If the affinity requirements specified by this field are not met at<br/>scheduling time, the pod will not be scheduled onto the node.<br/>If the affinity requirements specified by this field cease to be met<br/>at some point during pod execution (e.g. due to an update), the system<br/>may or may not try to eventually evict the pod from its node.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [nodeSelectorTerms](#lmdeploymentspecollamaaffinitynodeaffinityrequiredduringschedulingignoredduringexecutionnodeselectortermsindex) | []object | Required. A list of node selector terms. The terms are ORed. | true |
### LMDeployment.spec.ollama.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[index]

A null or empty node selector term matches no objects. The requirements of<br/>them are ANDed.<br/>The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecollamaaffinitynodeaffinityrequiredduringschedulingignoredduringexecutionnodeselectortermsindexmatchexpressionsindex) | []object | A list of node selector requirements by node's labels. | false |
| [matchFields](#lmdeploymentspecollamaaffinitynodeaffinityrequiredduringschedulingignoredduringexecutionnodeselectortermsindexmatchfieldsindex) | []object | A list of node selector requirements by node's fields. | false |
### LMDeployment.spec.ollama.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[index].matchExpressions[index]

A node selector requirement is a selector that contains values, a key, and an operator<br/>that relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The label key that the selector applies to. | true |
| operator | string | Represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt. | true |
| values | []string | An array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. If the operator is Gt or Lt, the values<br/>array must have a single element, which will be interpreted as an integer.<br/>This array is replaced during a strategic merge patch. | false |
### LMDeployment.spec.ollama.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[index].matchFields[index]

A node selector requirement is a selector that contains values, a key, and an operator<br/>that relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The label key that the selector applies to. | true |
| operator | string | Represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt. | true |
| values | []string | An array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. If the operator is Gt or Lt, the values<br/>array must have a single element, which will be interpreted as an integer.<br/>This array is replaced during a strategic merge patch. | false |
### LMDeployment.spec.ollama.affinity.podAffinity

Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [preferredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspecollamaaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindex) | []object | The scheduler will prefer to schedule pods to nodes that satisfy<br/>the affinity expressions specified by this field, but it may choose<br/>a node that violates one or more of the expressions. The node that is<br/>most preferred is the one with the greatest sum of weights, i.e.<br/>for each node that meets all of the scheduling requirements (resource<br/>request, requiredDuringScheduling affinity expressions, etc.),<br/>compute a sum by iterating through the elements of this field and adding<br/>"weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the<br/>node(s) with the highest sum are the most preferred. | false |
| [requiredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspecollamaaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindex) | []object | If the affinity requirements specified by this field are not met at<br/>scheduling time, the pod will not be scheduled onto the node.<br/>If the affinity requirements specified by this field cease to be met<br/>at some point during pod execution (e.g. due to a pod label update), the<br/>system may or may not try to eventually evict the pod from its node.<br/>When there are multiple elements, the lists of nodes corresponding to each<br/>podAffinityTerm are intersected, i.e. all terms must be satisfied. | false |
### LMDeployment.spec.ollama.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index]

The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [podAffinityTerm](#lmdeploymentspecollamaaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinityterm) | object | Required. A pod affinity term, associated with the corresponding weight. | true |
| weight | integer | weight associated with matching the corresponding podAffinityTerm,<br/>in the range 1-100.<br/>*Format*: int32<br/> | true |
### LMDeployment.spec.ollama.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm

Required. A pod affinity term, associated with the corresponding weight.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| topologyKey | string | This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching<br/>the labelSelector in the specified namespaces, where co-located is defined as running on a node<br/>whose value of the label with key topologyKey matches that of any node on which any of the<br/>selected pods is running.<br/>Empty topologyKey is not allowed. | true |
| [labelSelector](#lmdeploymentspecollamaaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselector) | object | A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods. | false |
| matchLabelKeys | []string | MatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both matchLabelKeys and labelSelector.<br/>Also, matchLabelKeys cannot be set when labelSelector isn't set. | false |
| mismatchLabelKeys | []string | MismatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.<br/>Also, mismatchLabelKeys cannot be set when labelSelector isn't set. | false |
| [namespaceSelector](#lmdeploymentspecollamaaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselector) | object | A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces. | false |
| namespaces | []string | namespaces specifies a static list of namespace names that the term applies to.<br/>The term is applied to the union of the namespaces listed in this field<br/>and the ones selected by namespaceSelector.<br/>null or empty namespaces list and null namespaceSelector means "this pod's namespace". | false |
### LMDeployment.spec.ollama.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector

A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecollamaaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.ollama.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.ollama.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector

A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecollamaaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.ollama.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.ollama.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index]

Defines a set of pods (namely those matching the labelSelector<br/>relative to the given namespace(s)) that this pod should be<br/>co-located (affinity) or not co-located (anti-affinity) with,<br/>where co-located is defined as running on a node whose value of<br/>the label with key &lt;topologyKey&gt; matches that of any node on which<br/>a pod of the set of pods is running

| Name | Type | Description | Required |
|------|------|-------------|----------|
| topologyKey | string | This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching<br/>the labelSelector in the specified namespaces, where co-located is defined as running on a node<br/>whose value of the label with key topologyKey matches that of any node on which any of the<br/>selected pods is running.<br/>Empty topologyKey is not allowed. | true |
| [labelSelector](#lmdeploymentspecollamaaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexlabelselector) | object | A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods. | false |
| matchLabelKeys | []string | MatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both matchLabelKeys and labelSelector.<br/>Also, matchLabelKeys cannot be set when labelSelector isn't set. | false |
| mismatchLabelKeys | []string | MismatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.<br/>Also, mismatchLabelKeys cannot be set when labelSelector isn't set. | false |
| [namespaceSelector](#lmdeploymentspecollamaaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselector) | object | A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces. | false |
| namespaces | []string | namespaces specifies a static list of namespace names that the term applies to.<br/>The term is applied to the union of the namespaces listed in this field<br/>and the ones selected by namespaceSelector.<br/>null or empty namespaces list and null namespaceSelector means "this pod's namespace". | false |
### LMDeployment.spec.ollama.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector

A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecollamaaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexlabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.ollama.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.ollama.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector

A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecollamaaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.ollama.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.ollama.affinity.podAntiAffinity

Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [preferredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspecollamaaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindex) | []object | The scheduler will prefer to schedule pods to nodes that satisfy<br/>the anti-affinity expressions specified by this field, but it may choose<br/>a node that violates one or more of the expressions. The node that is<br/>most preferred is the one with the greatest sum of weights, i.e.<br/>for each node that meets all of the scheduling requirements (resource<br/>request, requiredDuringScheduling anti-affinity expressions, etc.),<br/>compute a sum by iterating through the elements of this field and adding<br/>"weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the<br/>node(s) with the highest sum are the most preferred. | false |
| [requiredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspecollamaaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindex) | []object | If the anti-affinity requirements specified by this field are not met at<br/>scheduling time, the pod will not be scheduled onto the node.<br/>If the anti-affinity requirements specified by this field cease to be met<br/>at some point during pod execution (e.g. due to a pod label update), the<br/>system may or may not try to eventually evict the pod from its node.<br/>When there are multiple elements, the lists of nodes corresponding to each<br/>podAffinityTerm are intersected, i.e. all terms must be satisfied. | false |
### LMDeployment.spec.ollama.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index]

The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [podAffinityTerm](#lmdeploymentspecollamaaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinityterm) | object | Required. A pod affinity term, associated with the corresponding weight. | true |
| weight | integer | weight associated with matching the corresponding podAffinityTerm,<br/>in the range 1-100.<br/>*Format*: int32<br/> | true |
### LMDeployment.spec.ollama.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm

Required. A pod affinity term, associated with the corresponding weight.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| topologyKey | string | This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching<br/>the labelSelector in the specified namespaces, where co-located is defined as running on a node<br/>whose value of the label with key topologyKey matches that of any node on which any of the<br/>selected pods is running.<br/>Empty topologyKey is not allowed. | true |
| [labelSelector](#lmdeploymentspecollamaaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselector) | object | A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods. | false |
| matchLabelKeys | []string | MatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both matchLabelKeys and labelSelector.<br/>Also, matchLabelKeys cannot be set when labelSelector isn't set. | false |
| mismatchLabelKeys | []string | MismatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.<br/>Also, mismatchLabelKeys cannot be set when labelSelector isn't set. | false |
| [namespaceSelector](#lmdeploymentspecollamaaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselector) | object | A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces. | false |
| namespaces | []string | namespaces specifies a static list of namespace names that the term applies to.<br/>The term is applied to the union of the namespaces listed in this field<br/>and the ones selected by namespaceSelector.<br/>null or empty namespaces list and null namespaceSelector means "this pod's namespace". | false |
### LMDeployment.spec.ollama.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector

A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecollamaaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.ollama.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.ollama.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector

A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecollamaaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.ollama.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.ollama.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index]

Defines a set of pods (namely those matching the labelSelector<br/>relative to the given namespace(s)) that this pod should be<br/>co-located (affinity) or not co-located (anti-affinity) with,<br/>where co-located is defined as running on a node whose value of<br/>the label with key &lt;topologyKey&gt; matches that of any node on which<br/>a pod of the set of pods is running

| Name | Type | Description | Required |
|------|------|-------------|----------|
| topologyKey | string | This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching<br/>the labelSelector in the specified namespaces, where co-located is defined as running on a node<br/>whose value of the label with key topologyKey matches that of any node on which any of the<br/>selected pods is running.<br/>Empty topologyKey is not allowed. | true |
| [labelSelector](#lmdeploymentspecollamaaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexlabelselector) | object | A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods. | false |
| matchLabelKeys | []string | MatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both matchLabelKeys and labelSelector.<br/>Also, matchLabelKeys cannot be set when labelSelector isn't set. | false |
| mismatchLabelKeys | []string | MismatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.<br/>Also, mismatchLabelKeys cannot be set when labelSelector isn't set. | false |
| [namespaceSelector](#lmdeploymentspecollamaaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselector) | object | A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces. | false |
| namespaces | []string | namespaces specifies a static list of namespace names that the term applies to.<br/>The term is applied to the union of the namespaces listed in this field<br/>and the ones selected by namespaceSelector.<br/>null or empty namespaces list and null namespaceSelector means "this pod's namespace". | false |
### LMDeployment.spec.ollama.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector

A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecollamaaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexlabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.ollama.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.ollama.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector

A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecollamaaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.ollama.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.ollama.resources

Resources defines the resource requirements for Ollama pods

| Name | Type | Description | Required |
|------|------|-------------|----------|
| limits | map[string]int or string | Limits describes the maximum amount of compute resources allowed | false |
| requests | map[string]int or string | Requests describes the minimum amount of compute resources required | false |
### LMDeployment.spec.ollama.service

Service defines the service configuration for Ollama

| Name | Type | Description | Required |
|------|------|-------------|----------|
| port | integer | Port is the port to expose the service<br/>*Format*: int32<br/>*Minimum*: 1<br/>*Maximum*: 65535<br/> | false |
| type | enum | Type is the type of service to expose<br/>*Enum*: ClusterIP, NodePort, LoadBalancer<br/> | false |
### LMDeployment.spec.openwebui

OpenWebUI defines the OpenWebUI deployment configuration

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [affinity](#lmdeploymentspecopenwebuiaffinity) | object | Affinity defines pod affinity and anti-affinity rules for OpenWebUI pods | false |
| enabled | boolean | Enabled determines if OpenWebUI should be deployed | false |
| [envVars](#lmdeploymentspecopenwebuienvvarsindex) | []object | EnvVars defines environment variables for the Pipelines | false |
| image | string | Image is the OpenWebUI container image to use (including tag) | false |
| [ingress](#lmdeploymentspecopenwebuiingress) | object | Ingress defines the ingress configuration for OpenWebUI | false |
| [langfuse](#lmdeploymentspecopenwebuilangfuse) | object | Langfuse defines the Langfuse monitoring configuration for OpenWebUI | false |
| [persistence](#lmdeploymentspecopenwebuipersistence) | object | Persistence defines OpenWebUI persistence configuration | false |
| [pipelines](#lmdeploymentspecopenwebuipipelines) | object | Pipelines defines the OpenWebUI Pipelines configuration | false |
| [redis](#lmdeploymentspecopenwebuiredis) | object | Redis defines the Redis configuration for OpenWebUI | false |
| replicas | integer | Replicas is the number of OpenWebUI pods to run<br/>*Format*: int32<br/>*Minimum*: 1<br/>*Maximum*: 5<br/> | false |
| [resources](#lmdeploymentspecopenwebuiresources) | object | Resources defines the resource requirements for OpenWebUI pods | false |
| [service](#lmdeploymentspecopenwebuiservice) | object | Service defines the service configuration for OpenWebUI | false |
### LMDeployment.spec.openwebui.affinity

Affinity defines pod affinity and anti-affinity rules for OpenWebUI pods

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [nodeAffinity](#lmdeploymentspecopenwebuiaffinitynodeaffinity) | object | Describes node affinity scheduling rules for the pod. | false |
| [podAffinity](#lmdeploymentspecopenwebuiaffinitypodaffinity) | object | Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)). | false |
| [podAntiAffinity](#lmdeploymentspecopenwebuiaffinitypodantiaffinity) | object | Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)). | false |
### LMDeployment.spec.openwebui.affinity.nodeAffinity

Describes node affinity scheduling rules for the pod.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [preferredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspecopenwebuiaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindex) | []object | The scheduler will prefer to schedule pods to nodes that satisfy<br/>the affinity expressions specified by this field, but it may choose<br/>a node that violates one or more of the expressions. The node that is<br/>most preferred is the one with the greatest sum of weights, i.e.<br/>for each node that meets all of the scheduling requirements (resource<br/>request, requiredDuringScheduling affinity expressions, etc.),<br/>compute a sum by iterating through the elements of this field and adding<br/>"weight" to the sum if the node matches the corresponding matchExpressions; the<br/>node(s) with the highest sum are the most preferred. | false |
| [requiredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspecopenwebuiaffinitynodeaffinityrequiredduringschedulingignoredduringexecution) | object | If the affinity requirements specified by this field are not met at<br/>scheduling time, the pod will not be scheduled onto the node.<br/>If the affinity requirements specified by this field cease to be met<br/>at some point during pod execution (e.g. due to an update), the system<br/>may or may not try to eventually evict the pod from its node. | false |
### LMDeployment.spec.openwebui.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index]

An empty preferred scheduling term matches all objects with implicit weight 0<br/>(i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [preference](#lmdeploymentspecopenwebuiaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindexpreference) | object | A node selector term, associated with the corresponding weight. | true |
| weight | integer | Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.<br/>*Format*: int32<br/> | true |
### LMDeployment.spec.openwebui.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].preference

A node selector term, associated with the corresponding weight.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecopenwebuiaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindexpreferencematchexpressionsindex) | []object | A list of node selector requirements by node's labels. | false |
| [matchFields](#lmdeploymentspecopenwebuiaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindexpreferencematchfieldsindex) | []object | A list of node selector requirements by node's fields. | false |
### LMDeployment.spec.openwebui.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].preference.matchExpressions[index]

A node selector requirement is a selector that contains values, a key, and an operator<br/>that relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The label key that the selector applies to. | true |
| operator | string | Represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt. | true |
| values | []string | An array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. If the operator is Gt or Lt, the values<br/>array must have a single element, which will be interpreted as an integer.<br/>This array is replaced during a strategic merge patch. | false |
### LMDeployment.spec.openwebui.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].preference.matchFields[index]

A node selector requirement is a selector that contains values, a key, and an operator<br/>that relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The label key that the selector applies to. | true |
| operator | string | Represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt. | true |
| values | []string | An array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. If the operator is Gt or Lt, the values<br/>array must have a single element, which will be interpreted as an integer.<br/>This array is replaced during a strategic merge patch. | false |
### LMDeployment.spec.openwebui.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution

If the affinity requirements specified by this field are not met at<br/>scheduling time, the pod will not be scheduled onto the node.<br/>If the affinity requirements specified by this field cease to be met<br/>at some point during pod execution (e.g. due to an update), the system<br/>may or may not try to eventually evict the pod from its node.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [nodeSelectorTerms](#lmdeploymentspecopenwebuiaffinitynodeaffinityrequiredduringschedulingignoredduringexecutionnodeselectortermsindex) | []object | Required. A list of node selector terms. The terms are ORed. | true |
### LMDeployment.spec.openwebui.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[index]

A null or empty node selector term matches no objects. The requirements of<br/>them are ANDed.<br/>The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecopenwebuiaffinitynodeaffinityrequiredduringschedulingignoredduringexecutionnodeselectortermsindexmatchexpressionsindex) | []object | A list of node selector requirements by node's labels. | false |
| [matchFields](#lmdeploymentspecopenwebuiaffinitynodeaffinityrequiredduringschedulingignoredduringexecutionnodeselectortermsindexmatchfieldsindex) | []object | A list of node selector requirements by node's fields. | false |
### LMDeployment.spec.openwebui.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[index].matchExpressions[index]

A node selector requirement is a selector that contains values, a key, and an operator<br/>that relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The label key that the selector applies to. | true |
| operator | string | Represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt. | true |
| values | []string | An array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. If the operator is Gt or Lt, the values<br/>array must have a single element, which will be interpreted as an integer.<br/>This array is replaced during a strategic merge patch. | false |
### LMDeployment.spec.openwebui.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[index].matchFields[index]

A node selector requirement is a selector that contains values, a key, and an operator<br/>that relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The label key that the selector applies to. | true |
| operator | string | Represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt. | true |
| values | []string | An array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. If the operator is Gt or Lt, the values<br/>array must have a single element, which will be interpreted as an integer.<br/>This array is replaced during a strategic merge patch. | false |
### LMDeployment.spec.openwebui.affinity.podAffinity

Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [preferredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspecopenwebuiaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindex) | []object | The scheduler will prefer to schedule pods to nodes that satisfy<br/>the affinity expressions specified by this field, but it may choose<br/>a node that violates one or more of the expressions. The node that is<br/>most preferred is the one with the greatest sum of weights, i.e.<br/>for each node that meets all of the scheduling requirements (resource<br/>request, requiredDuringScheduling affinity expressions, etc.),<br/>compute a sum by iterating through the elements of this field and adding<br/>"weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the<br/>node(s) with the highest sum are the most preferred. | false |
| [requiredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspecopenwebuiaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindex) | []object | If the affinity requirements specified by this field are not met at<br/>scheduling time, the pod will not be scheduled onto the node.<br/>If the affinity requirements specified by this field cease to be met<br/>at some point during pod execution (e.g. due to a pod label update), the<br/>system may or may not try to eventually evict the pod from its node.<br/>When there are multiple elements, the lists of nodes corresponding to each<br/>podAffinityTerm are intersected, i.e. all terms must be satisfied. | false |
### LMDeployment.spec.openwebui.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index]

The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [podAffinityTerm](#lmdeploymentspecopenwebuiaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinityterm) | object | Required. A pod affinity term, associated with the corresponding weight. | true |
| weight | integer | weight associated with matching the corresponding podAffinityTerm,<br/>in the range 1-100.<br/>*Format*: int32<br/> | true |
### LMDeployment.spec.openwebui.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm

Required. A pod affinity term, associated with the corresponding weight.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| topologyKey | string | This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching<br/>the labelSelector in the specified namespaces, where co-located is defined as running on a node<br/>whose value of the label with key topologyKey matches that of any node on which any of the<br/>selected pods is running.<br/>Empty topologyKey is not allowed. | true |
| [labelSelector](#lmdeploymentspecopenwebuiaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselector) | object | A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods. | false |
| matchLabelKeys | []string | MatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both matchLabelKeys and labelSelector.<br/>Also, matchLabelKeys cannot be set when labelSelector isn't set. | false |
| mismatchLabelKeys | []string | MismatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.<br/>Also, mismatchLabelKeys cannot be set when labelSelector isn't set. | false |
| [namespaceSelector](#lmdeploymentspecopenwebuiaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselector) | object | A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces. | false |
| namespaces | []string | namespaces specifies a static list of namespace names that the term applies to.<br/>The term is applied to the union of the namespaces listed in this field<br/>and the ones selected by namespaceSelector.<br/>null or empty namespaces list and null namespaceSelector means "this pod's namespace". | false |
### LMDeployment.spec.openwebui.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector

A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecopenwebuiaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.openwebui.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.openwebui.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector

A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecopenwebuiaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.openwebui.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.openwebui.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index]

Defines a set of pods (namely those matching the labelSelector<br/>relative to the given namespace(s)) that this pod should be<br/>co-located (affinity) or not co-located (anti-affinity) with,<br/>where co-located is defined as running on a node whose value of<br/>the label with key &lt;topologyKey&gt; matches that of any node on which<br/>a pod of the set of pods is running

| Name | Type | Description | Required |
|------|------|-------------|----------|
| topologyKey | string | This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching<br/>the labelSelector in the specified namespaces, where co-located is defined as running on a node<br/>whose value of the label with key topologyKey matches that of any node on which any of the<br/>selected pods is running.<br/>Empty topologyKey is not allowed. | true |
| [labelSelector](#lmdeploymentspecopenwebuiaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexlabelselector) | object | A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods. | false |
| matchLabelKeys | []string | MatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both matchLabelKeys and labelSelector.<br/>Also, matchLabelKeys cannot be set when labelSelector isn't set. | false |
| mismatchLabelKeys | []string | MismatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.<br/>Also, mismatchLabelKeys cannot be set when labelSelector isn't set. | false |
| [namespaceSelector](#lmdeploymentspecopenwebuiaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselector) | object | A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces. | false |
| namespaces | []string | namespaces specifies a static list of namespace names that the term applies to.<br/>The term is applied to the union of the namespaces listed in this field<br/>and the ones selected by namespaceSelector.<br/>null or empty namespaces list and null namespaceSelector means "this pod's namespace". | false |
### LMDeployment.spec.openwebui.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector

A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecopenwebuiaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexlabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.openwebui.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.openwebui.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector

A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecopenwebuiaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.openwebui.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.openwebui.affinity.podAntiAffinity

Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [preferredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspecopenwebuiaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindex) | []object | The scheduler will prefer to schedule pods to nodes that satisfy<br/>the anti-affinity expressions specified by this field, but it may choose<br/>a node that violates one or more of the expressions. The node that is<br/>most preferred is the one with the greatest sum of weights, i.e.<br/>for each node that meets all of the scheduling requirements (resource<br/>request, requiredDuringScheduling anti-affinity expressions, etc.),<br/>compute a sum by iterating through the elements of this field and adding<br/>"weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the<br/>node(s) with the highest sum are the most preferred. | false |
| [requiredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspecopenwebuiaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindex) | []object | If the anti-affinity requirements specified by this field are not met at<br/>scheduling time, the pod will not be scheduled onto the node.<br/>If the anti-affinity requirements specified by this field cease to be met<br/>at some point during pod execution (e.g. due to a pod label update), the<br/>system may or may not try to eventually evict the pod from its node.<br/>When there are multiple elements, the lists of nodes corresponding to each<br/>podAffinityTerm are intersected, i.e. all terms must be satisfied. | false |
### LMDeployment.spec.openwebui.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index]

The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [podAffinityTerm](#lmdeploymentspecopenwebuiaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinityterm) | object | Required. A pod affinity term, associated with the corresponding weight. | true |
| weight | integer | weight associated with matching the corresponding podAffinityTerm,<br/>in the range 1-100.<br/>*Format*: int32<br/> | true |
### LMDeployment.spec.openwebui.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm

Required. A pod affinity term, associated with the corresponding weight.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| topologyKey | string | This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching<br/>the labelSelector in the specified namespaces, where co-located is defined as running on a node<br/>whose value of the label with key topologyKey matches that of any node on which any of the<br/>selected pods is running.<br/>Empty topologyKey is not allowed. | true |
| [labelSelector](#lmdeploymentspecopenwebuiaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselector) | object | A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods. | false |
| matchLabelKeys | []string | MatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both matchLabelKeys and labelSelector.<br/>Also, matchLabelKeys cannot be set when labelSelector isn't set. | false |
| mismatchLabelKeys | []string | MismatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.<br/>Also, mismatchLabelKeys cannot be set when labelSelector isn't set. | false |
| [namespaceSelector](#lmdeploymentspecopenwebuiaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselector) | object | A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces. | false |
| namespaces | []string | namespaces specifies a static list of namespace names that the term applies to.<br/>The term is applied to the union of the namespaces listed in this field<br/>and the ones selected by namespaceSelector.<br/>null or empty namespaces list and null namespaceSelector means "this pod's namespace". | false |
### LMDeployment.spec.openwebui.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector

A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecopenwebuiaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.openwebui.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.openwebui.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector

A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecopenwebuiaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.openwebui.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.openwebui.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index]

Defines a set of pods (namely those matching the labelSelector<br/>relative to the given namespace(s)) that this pod should be<br/>co-located (affinity) or not co-located (anti-affinity) with,<br/>where co-located is defined as running on a node whose value of<br/>the label with key &lt;topologyKey&gt; matches that of any node on which<br/>a pod of the set of pods is running

| Name | Type | Description | Required |
|------|------|-------------|----------|
| topologyKey | string | This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching<br/>the labelSelector in the specified namespaces, where co-located is defined as running on a node<br/>whose value of the label with key topologyKey matches that of any node on which any of the<br/>selected pods is running.<br/>Empty topologyKey is not allowed. | true |
| [labelSelector](#lmdeploymentspecopenwebuiaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexlabelselector) | object | A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods. | false |
| matchLabelKeys | []string | MatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both matchLabelKeys and labelSelector.<br/>Also, matchLabelKeys cannot be set when labelSelector isn't set. | false |
| mismatchLabelKeys | []string | MismatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.<br/>Also, mismatchLabelKeys cannot be set when labelSelector isn't set. | false |
| [namespaceSelector](#lmdeploymentspecopenwebuiaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselector) | object | A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces. | false |
| namespaces | []string | namespaces specifies a static list of namespace names that the term applies to.<br/>The term is applied to the union of the namespaces listed in this field<br/>and the ones selected by namespaceSelector.<br/>null or empty namespaces list and null namespaceSelector means "this pod's namespace". | false |
### LMDeployment.spec.openwebui.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector

A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecopenwebuiaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexlabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.openwebui.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.openwebui.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector

A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecopenwebuiaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.openwebui.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.openwebui.envVars[index]

EnvVar represents an environment variable present in a Container.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the environment variable. Must be a C_IDENTIFIER. | true |
| value | string | Variable references $(VAR_NAME) are expanded<br/>using the previously defined environment variables in the container and<br/>any service environment variables. If a variable cannot be resolved,<br/>the reference in the input string will be unchanged. Double $$ are reduced<br/>to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.<br/>"$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)".<br/>Escaped references will never be expanded, regardless of whether the variable<br/>exists or not.<br/>Defaults to "". | false |
| [valueFrom](#lmdeploymentspecopenwebuienvvarsindexvaluefrom) | object | Source for the environment variable's value. Cannot be used if value is not empty. | false |
### LMDeployment.spec.openwebui.envVars[index].valueFrom

Source for the environment variable's value. Cannot be used if value is not empty.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [configMapKeyRef](#lmdeploymentspecopenwebuienvvarsindexvaluefromconfigmapkeyref) | object | Selects a key of a ConfigMap. | false |
| [fieldRef](#lmdeploymentspecopenwebuienvvarsindexvaluefromfieldref) | object | Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['&lt;KEY&gt;']`, `metadata.annotations['&lt;KEY&gt;']`,<br/>spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs. | false |
| [resourceFieldRef](#lmdeploymentspecopenwebuienvvarsindexvaluefromresourcefieldref) | object | Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported. | false |
| [secretKeyRef](#lmdeploymentspecopenwebuienvvarsindexvaluefromsecretkeyref) | object | Selects a key of a secret in the pod's namespace | false |
### LMDeployment.spec.openwebui.envVars[index].valueFrom.configMapKeyRef

Selects a key of a ConfigMap.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The key to select. | true |
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
| optional | boolean | Specify whether the ConfigMap or its key must be defined | false |
### LMDeployment.spec.openwebui.envVars[index].valueFrom.fieldRef

Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['&lt;KEY&gt;']`, `metadata.annotations['&lt;KEY&gt;']`,<br/>spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| fieldPath | string | Path of the field to select in the specified API version. | true |
| apiVersion | string | Version of the schema the FieldPath is written in terms of, defaults to "v1". | false |
### LMDeployment.spec.openwebui.envVars[index].valueFrom.resourceFieldRef

Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| resource | string | Required: resource to select | true |
| containerName | string | Container name: required for volumes, optional for env vars | false |
| divisor | int or string | Specifies the output format of the exposed resources, defaults to "1" | false |
### LMDeployment.spec.openwebui.envVars[index].valueFrom.secretKeyRef

Selects a key of a secret in the pod's namespace

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The key of the secret to select from.  Must be a valid secret key. | true |
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
| optional | boolean | Specify whether the Secret or its key must be defined | false |
### LMDeployment.spec.openwebui.ingress

Ingress defines the ingress configuration for OpenWebUI

| Name | Type | Description | Required |
|------|------|-------------|----------|
| annotations | map[string]string | Annotations are custom annotations for the Ingress | false |
| host | string | Host is the hostname for the Ingress | false |
### LMDeployment.spec.openwebui.langfuse

Langfuse defines the Langfuse monitoring configuration for OpenWebUI

| Name | Type | Description | Required |
|------|------|-------------|----------|
| debug | boolean | Debug enables debug logging for Langfuse | false |
| enabled | boolean | Enabled determines if Langfuse monitoring should be enabled | false |
| environment | string | Environment is the environment name (e.g., "production", "staging", "development") | false |
| projectName | string | ProjectName is the name of the project for Langfuse<br/>If not provided, will default to deployment name | false |
| [secretRef](#lmdeploymentspecopenwebuilangfusesecretref) | object | SecretRef is the reference to a Kubernetes secret containing Langfuse credentials<br/>The secret should contain: LANGFUSE_PUBLIC_KEY, LANGFUSE_SECRET_KEY | false |
| url | string | URL is the Langfuse server URL<br/>Format: https://cloud.langfuse.com or http://localhost:3000 | false |
### LMDeployment.spec.openwebui.langfuse.secretRef

SecretRef is the reference to a Kubernetes secret containing Langfuse credentials<br/>The secret should contain: LANGFUSE_PUBLIC_KEY, LANGFUSE_SECRET_KEY

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | name is unique within a namespace to reference a secret resource. | false |
| namespace | string | namespace defines the space within which the secret name must be unique. | false |
### LMDeployment.spec.openwebui.persistence

Persistence defines OpenWebUI persistence configuration

| Name | Type | Description | Required |
|------|------|-------------|----------|
| enabled | boolean | Enabled determines if OpenWebUI data should be persisted | false |
| size | string | Size is the size of the persistent volume | false |
| storageClass | string | StorageClass is the storage class to use for persistent volumes | false |
### LMDeployment.spec.openwebui.pipelines

Pipelines defines the OpenWebUI Pipelines configuration

| Name | Type | Description | Required |
|------|------|-------------|----------|
| enabled | boolean | Enabled determines if OpenWebUI Pipelines should be deployed | false |
| [envVars](#lmdeploymentspecopenwebuipipelinesenvvarsindex) | []object | EnvVars defines environment variables for the Pipelines | false |
| image | string | Image is the Pipelines container image to use (including tag) | false |
| [persistence](#lmdeploymentspecopenwebuipipelinespersistence) | object | Persistence defines Pipelines persistence configuration | false |
| pipelineUrls | []string | PipelineURLs is a list of URLs to fetch pipeline definitions from<br/>Format: https://github.com/open-webui/pipelines/blob/main/examples/filters/example.py | false |
| pipelinesDir | string | PipelinesDir is the directory containing pipeline definitions<br/>Default: /app/pipelines | false |
| port | integer | Port is the port the Pipelines service exposes<br/>*Format*: int32<br/>*Minimum*: 1<br/>*Maximum*: 65535<br/> | false |
| replicas | integer | Replicas is the number of Pipelines pods to run<br/>*Format*: int32<br/>*Minimum*: 1<br/>*Maximum*: 3<br/> | false |
| [resources](#lmdeploymentspecopenwebuipipelinesresources) | object | Resources defines the resource requirements for Pipelines pods | false |
| [service](#lmdeploymentspecopenwebuipipelinesservice) | object | Service describes service to expose the Pipelines | false |
| [volumeMounts](#lmdeploymentspecopenwebuipipelinesvolumemountsindex) | []object | VolumeMounts defines volume mounts for the Pipelines | false |
| [volumes](#lmdeploymentspecopenwebuipipelinesvolumesindex) | []object | Volumes defines volumes for the Pipelines | false |
### LMDeployment.spec.openwebui.pipelines.envVars[index]

EnvVar represents an environment variable present in a Container.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the environment variable. Must be a C_IDENTIFIER. | true |
| value | string | Variable references $(VAR_NAME) are expanded<br/>using the previously defined environment variables in the container and<br/>any service environment variables. If a variable cannot be resolved,<br/>the reference in the input string will be unchanged. Double $$ are reduced<br/>to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.<br/>"$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)".<br/>Escaped references will never be expanded, regardless of whether the variable<br/>exists or not.<br/>Defaults to "". | false |
| [valueFrom](#lmdeploymentspecopenwebuipipelinesenvvarsindexvaluefrom) | object | Source for the environment variable's value. Cannot be used if value is not empty. | false |
### LMDeployment.spec.openwebui.pipelines.envVars[index].valueFrom

Source for the environment variable's value. Cannot be used if value is not empty.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [configMapKeyRef](#lmdeploymentspecopenwebuipipelinesenvvarsindexvaluefromconfigmapkeyref) | object | Selects a key of a ConfigMap. | false |
| [fieldRef](#lmdeploymentspecopenwebuipipelinesenvvarsindexvaluefromfieldref) | object | Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['&lt;KEY&gt;']`, `metadata.annotations['&lt;KEY&gt;']`,<br/>spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs. | false |
| [resourceFieldRef](#lmdeploymentspecopenwebuipipelinesenvvarsindexvaluefromresourcefieldref) | object | Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported. | false |
| [secretKeyRef](#lmdeploymentspecopenwebuipipelinesenvvarsindexvaluefromsecretkeyref) | object | Selects a key of a secret in the pod's namespace | false |
### LMDeployment.spec.openwebui.pipelines.envVars[index].valueFrom.configMapKeyRef

Selects a key of a ConfigMap.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The key to select. | true |
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
| optional | boolean | Specify whether the ConfigMap or its key must be defined | false |
### LMDeployment.spec.openwebui.pipelines.envVars[index].valueFrom.fieldRef

Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['&lt;KEY&gt;']`, `metadata.annotations['&lt;KEY&gt;']`,<br/>spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| fieldPath | string | Path of the field to select in the specified API version. | true |
| apiVersion | string | Version of the schema the FieldPath is written in terms of, defaults to "v1". | false |
### LMDeployment.spec.openwebui.pipelines.envVars[index].valueFrom.resourceFieldRef

Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| resource | string | Required: resource to select | true |
| containerName | string | Container name: required for volumes, optional for env vars | false |
| divisor | int or string | Specifies the output format of the exposed resources, defaults to "1" | false |
### LMDeployment.spec.openwebui.pipelines.envVars[index].valueFrom.secretKeyRef

Selects a key of a secret in the pod's namespace

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The key of the secret to select from.  Must be a valid secret key. | true |
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
| optional | boolean | Specify whether the Secret or its key must be defined | false |
### LMDeployment.spec.openwebui.pipelines.persistence

Persistence defines Pipelines persistence configuration

| Name | Type | Description | Required |
|------|------|-------------|----------|
| enabled | boolean | Enabled determines if Pipelines data should be persisted | false |
| size | string | Size is the size of the persistent volume | false |
| storageClass | string | StorageClass is the storage class to use for persistent volumes | false |
### LMDeployment.spec.openwebui.pipelines.resources

Resources defines the resource requirements for Pipelines pods

| Name | Type | Description | Required |
|------|------|-------------|----------|
| limits | map[string]int or string | Limits describes the maximum amount of compute resources allowed | false |
| requests | map[string]int or string | Requests describes the minimum amount of compute resources required | false |
### LMDeployment.spec.openwebui.pipelines.service

Service describes service to expose the Pipelines

| Name | Type | Description | Required |
|------|------|-------------|----------|
| port | integer | Port is the port to expose the service<br/>*Format*: int32<br/>*Minimum*: 1<br/>*Maximum*: 65535<br/> | false |
| type | enum | Type is the type of service to expose<br/>*Enum*: ClusterIP, NodePort, LoadBalancer<br/> | false |
### LMDeployment.spec.openwebui.pipelines.volumeMounts[index]

VolumeMount describes a mounting of a Volume within a container.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| mountPath | string | Path within the container at which the volume should be mounted.  Must<br/>not contain ':'. | true |
| name | string | This must match the Name of a Volume. | true |
| mountPropagation | string | mountPropagation determines how mounts are propagated from the host<br/>to container and the other way around.<br/>When not set, MountPropagationNone is used.<br/>This field is beta in 1.10.<br/>When RecursiveReadOnly is set to IfPossible or to Enabled, MountPropagation must be None or unspecified<br/>(which defaults to None). | false |
| readOnly | boolean | Mounted read-only if true, read-write otherwise (false or unspecified).<br/>Defaults to false. | false |
| recursiveReadOnly | string | RecursiveReadOnly specifies whether read-only mounts should be handled<br/>recursively.<br/><br/>If ReadOnly is false, this field has no meaning and must be unspecified.<br/><br/>If ReadOnly is true, and this field is set to Disabled, the mount is not made<br/>recursively read-only.  If this field is set to IfPossible, the mount is made<br/>recursively read-only, if it is supported by the container runtime.  If this<br/>field is set to Enabled, the mount is made recursively read-only if it is<br/>supported by the container runtime, otherwise the pod will not be started and<br/>an error will be generated to indicate the reason.<br/><br/>If this field is set to IfPossible or Enabled, MountPropagation must be set to<br/>None (or be unspecified, which defaults to None).<br/><br/>If this field is not specified, it is treated as an equivalent of Disabled. | false |
| subPath | string | Path within the volume from which the container's volume should be mounted.<br/>Defaults to "" (volume's root). | false |
| subPathExpr | string | Expanded path within the volume from which the container's volume should be mounted.<br/>Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.<br/>Defaults to "" (volume's root).<br/>SubPathExpr and SubPath are mutually exclusive. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index]

Volume represents a named volume in a pod that may be accessed by any container in the pod.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | name of the volume.<br/>Must be a DNS_LABEL and unique within the pod.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names | true |
| [awsElasticBlockStore](#lmdeploymentspecopenwebuipipelinesvolumesindexawselasticblockstore) | object | awsElasticBlockStore represents an AWS Disk resource that is attached to a<br/>kubelet's host machine and then exposed to the pod.<br/>Deprecated: AWSElasticBlockStore is deprecated. All operations for the in-tree<br/>awsElasticBlockStore type are redirected to the ebs.csi.aws.com CSI driver.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore | false |
| [azureDisk](#lmdeploymentspecopenwebuipipelinesvolumesindexazuredisk) | object | azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.<br/>Deprecated: AzureDisk is deprecated. All operations for the in-tree azureDisk type<br/>are redirected to the disk.csi.azure.com CSI driver. | false |
| [azureFile](#lmdeploymentspecopenwebuipipelinesvolumesindexazurefile) | object | azureFile represents an Azure File Service mount on the host and bind mount to the pod.<br/>Deprecated: AzureFile is deprecated. All operations for the in-tree azureFile type<br/>are redirected to the file.csi.azure.com CSI driver. | false |
| [cephfs](#lmdeploymentspecopenwebuipipelinesvolumesindexcephfs) | object | cephFS represents a Ceph FS mount on the host that shares a pod's lifetime.<br/>Deprecated: CephFS is deprecated and the in-tree cephfs type is no longer supported. | false |
| [cinder](#lmdeploymentspecopenwebuipipelinesvolumesindexcinder) | object | cinder represents a cinder volume attached and mounted on kubelets host machine.<br/>Deprecated: Cinder is deprecated. All operations for the in-tree cinder type<br/>are redirected to the cinder.csi.openstack.org CSI driver.<br/>More info: https://examples.k8s.io/mysql-cinder-pd/README.md | false |
| [configMap](#lmdeploymentspecopenwebuipipelinesvolumesindexconfigmap) | object | configMap represents a configMap that should populate this volume | false |
| [csi](#lmdeploymentspecopenwebuipipelinesvolumesindexcsi) | object | csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers. | false |
| [downwardAPI](#lmdeploymentspecopenwebuipipelinesvolumesindexdownwardapi) | object | downwardAPI represents downward API about the pod that should populate this volume | false |
| [emptyDir](#lmdeploymentspecopenwebuipipelinesvolumesindexemptydir) | object | emptyDir represents a temporary directory that shares a pod's lifetime.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir | false |
| [ephemeral](#lmdeploymentspecopenwebuipipelinesvolumesindexephemeral) | object | ephemeral represents a volume that is handled by a cluster storage driver.<br/>The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts,<br/>and deleted when the pod is removed.<br/><br/>Use this if:<br/>a) the volume is only needed while the pod runs,<br/>b) features of normal volumes like restoring from snapshot or capacity<br/>   tracking are needed,<br/>c) the storage driver is specified through a storage class, and<br/>d) the storage driver supports dynamic volume provisioning through<br/>   a PersistentVolumeClaim (see EphemeralVolumeSource for more<br/>   information on the connection between this volume type<br/>   and PersistentVolumeClaim).<br/><br/>Use PersistentVolumeClaim or one of the vendor-specific<br/>APIs for volumes that persist for longer than the lifecycle<br/>of an individual pod.<br/><br/>Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to<br/>be used that way - see the documentation of the driver for<br/>more information.<br/><br/>A pod can use both types of ephemeral volumes and<br/>persistent volumes at the same time. | false |
| [fc](#lmdeploymentspecopenwebuipipelinesvolumesindexfc) | object | fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod. | false |
| [flexVolume](#lmdeploymentspecopenwebuipipelinesvolumesindexflexvolume) | object | flexVolume represents a generic volume resource that is<br/>provisioned/attached using an exec based plugin.<br/>Deprecated: FlexVolume is deprecated. Consider using a CSIDriver instead. | false |
| [flocker](#lmdeploymentspecopenwebuipipelinesvolumesindexflocker) | object | flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running.<br/>Deprecated: Flocker is deprecated and the in-tree flocker type is no longer supported. | false |
| [gcePersistentDisk](#lmdeploymentspecopenwebuipipelinesvolumesindexgcepersistentdisk) | object | gcePersistentDisk represents a GCE Disk resource that is attached to a<br/>kubelet's host machine and then exposed to the pod.<br/>Deprecated: GCEPersistentDisk is deprecated. All operations for the in-tree<br/>gcePersistentDisk type are redirected to the pd.csi.storage.gke.io CSI driver.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk | false |
| [gitRepo](#lmdeploymentspecopenwebuipipelinesvolumesindexgitrepo) | object | gitRepo represents a git repository at a particular revision.<br/>Deprecated: GitRepo is deprecated. To provision a container with a git repo, mount an<br/>EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir<br/>into the Pod's container. | false |
| [glusterfs](#lmdeploymentspecopenwebuipipelinesvolumesindexglusterfs) | object | glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.<br/>Deprecated: Glusterfs is deprecated and the in-tree glusterfs type is no longer supported.<br/>More info: https://examples.k8s.io/volumes/glusterfs/README.md | false |
| [hostPath](#lmdeploymentspecopenwebuipipelinesvolumesindexhostpath) | object | hostPath represents a pre-existing file or directory on the host<br/>machine that is directly exposed to the container. This is generally<br/>used for system agents or other privileged things that are allowed<br/>to see the host machine. Most containers will NOT need this.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath | false |
| [image](#lmdeploymentspecopenwebuipipelinesvolumesindeximage) | object | image represents an OCI object (a container image or artifact) pulled and mounted on the kubelet's host machine.<br/>The volume is resolved at pod startup depending on which PullPolicy value is provided:<br/><br/>- Always: the kubelet always attempts to pull the reference. Container creation will fail If the pull fails.<br/>- Never: the kubelet never pulls the reference and only uses a local image or artifact. Container creation will fail if the reference isn't present.<br/>- IfNotPresent: the kubelet pulls if the reference isn't already present on disk. Container creation will fail if the reference isn't present and the pull fails.<br/><br/>The volume gets re-resolved if the pod gets deleted and recreated, which means that new remote content will become available on pod recreation.<br/>A failure to resolve or pull the image during pod startup will block containers from starting and may add significant latency. Failures will be retried using normal volume backoff and will be reported on the pod reason and message.<br/>The types of objects that may be mounted by this volume are defined by the container runtime implementation on a host machine and at minimum must include all valid types supported by the container image field.<br/>The OCI object gets mounted in a single directory (spec.containers[*].volumeMounts.mountPath) by merging the manifest layers in the same way as for container images.<br/>The volume will be mounted read-only (ro) and non-executable files (noexec).<br/>Sub path mounts for containers are not supported (spec.containers[*].volumeMounts.subpath) before 1.33.<br/>The field spec.securityContext.fsGroupChangePolicy has no effect on this volume type. | false |
| [iscsi](#lmdeploymentspecopenwebuipipelinesvolumesindexiscsi) | object | iscsi represents an ISCSI Disk resource that is attached to a<br/>kubelet's host machine and then exposed to the pod.<br/>More info: https://examples.k8s.io/volumes/iscsi/README.md | false |
| [nfs](#lmdeploymentspecopenwebuipipelinesvolumesindexnfs) | object | nfs represents an NFS mount on the host that shares a pod's lifetime<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs | false |
| [persistentVolumeClaim](#lmdeploymentspecopenwebuipipelinesvolumesindexpersistentvolumeclaim) | object | persistentVolumeClaimVolumeSource represents a reference to a<br/>PersistentVolumeClaim in the same namespace.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims | false |
| [photonPersistentDisk](#lmdeploymentspecopenwebuipipelinesvolumesindexphotonpersistentdisk) | object | photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine.<br/>Deprecated: PhotonPersistentDisk is deprecated and the in-tree photonPersistentDisk type is no longer supported. | false |
| [portworxVolume](#lmdeploymentspecopenwebuipipelinesvolumesindexportworxvolume) | object | portworxVolume represents a portworx volume attached and mounted on kubelets host machine.<br/>Deprecated: PortworxVolume is deprecated. All operations for the in-tree portworxVolume type<br/>are redirected to the pxd.portworx.com CSI driver when the CSIMigrationPortworx feature-gate<br/>is on. | false |
| [projected](#lmdeploymentspecopenwebuipipelinesvolumesindexprojected) | object | projected items for all in one resources secrets, configmaps, and downward API | false |
| [quobyte](#lmdeploymentspecopenwebuipipelinesvolumesindexquobyte) | object | quobyte represents a Quobyte mount on the host that shares a pod's lifetime.<br/>Deprecated: Quobyte is deprecated and the in-tree quobyte type is no longer supported. | false |
| [rbd](#lmdeploymentspecopenwebuipipelinesvolumesindexrbd) | object | rbd represents a Rados Block Device mount on the host that shares a pod's lifetime.<br/>Deprecated: RBD is deprecated and the in-tree rbd type is no longer supported.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md | false |
| [scaleIO](#lmdeploymentspecopenwebuipipelinesvolumesindexscaleio) | object | scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.<br/>Deprecated: ScaleIO is deprecated and the in-tree scaleIO type is no longer supported. | false |
| [secret](#lmdeploymentspecopenwebuipipelinesvolumesindexsecret) | object | secret represents a secret that should populate this volume.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#secret | false |
| [storageos](#lmdeploymentspecopenwebuipipelinesvolumesindexstorageos) | object | storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.<br/>Deprecated: StorageOS is deprecated and the in-tree storageos type is no longer supported. | false |
| [vsphereVolume](#lmdeploymentspecopenwebuipipelinesvolumesindexvspherevolume) | object | vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine.<br/>Deprecated: VsphereVolume is deprecated. All operations for the in-tree vsphereVolume type<br/>are redirected to the csi.vsphere.vmware.com CSI driver. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].awsElasticBlockStore

awsElasticBlockStore represents an AWS Disk resource that is attached to a<br/>kubelet's host machine and then exposed to the pod.<br/>Deprecated: AWSElasticBlockStore is deprecated. All operations for the in-tree<br/>awsElasticBlockStore type are redirected to the ebs.csi.aws.com CSI driver.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore

| Name | Type | Description | Required |
|------|------|-------------|----------|
| volumeID | string | volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume).<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore | true |
| fsType | string | fsType is the filesystem type of the volume that you want to mount.<br/>Tip: Ensure that the filesystem type is supported by the host operating system.<br/>Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore | false |
| partition | integer | partition is the partition in the volume that you want to mount.<br/>If omitted, the default is to mount by volume name.<br/>Examples: For volume /dev/sda1, you specify the partition as "1".<br/>Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).<br/>*Format*: int32<br/> | false |
| readOnly | boolean | readOnly value true will force the readOnly setting in VolumeMounts.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].azureDisk

azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.<br/>Deprecated: AzureDisk is deprecated. All operations for the in-tree azureDisk type<br/>are redirected to the disk.csi.azure.com CSI driver.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| diskName | string | diskName is the Name of the data disk in the blob storage | true |
| diskURI | string | diskURI is the URI of data disk in the blob storage | true |
| cachingMode | string | cachingMode is the Host Caching mode: None, Read Only, Read Write. | false |
| fsType | string | fsType is Filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br/>*Default*: ext4<br/> | false |
| kind | string | kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared | false |
| readOnly | boolean | readOnly Defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts.<br/>*Default*: false<br/> | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].azureFile

azureFile represents an Azure File Service mount on the host and bind mount to the pod.<br/>Deprecated: AzureFile is deprecated. All operations for the in-tree azureFile type<br/>are redirected to the file.csi.azure.com CSI driver.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| secretName | string | secretName is the  name of secret that contains Azure Storage Account Name and Key | true |
| shareName | string | shareName is the azure share Name | true |
| readOnly | boolean | readOnly defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].cephfs

cephFS represents a Ceph FS mount on the host that shares a pod's lifetime.<br/>Deprecated: CephFS is deprecated and the in-tree cephfs type is no longer supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| monitors | []string | monitors is Required: Monitors is a collection of Ceph monitors<br/>More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | true |
| path | string | path is Optional: Used as the mounted root, rather than the full Ceph tree, default is / | false |
| readOnly | boolean | readOnly is Optional: Defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts.<br/>More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | false |
| secretFile | string | secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret<br/>More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | false |
| [secretRef](#lmdeploymentspecopenwebuipipelinesvolumesindexcephfssecretref) | object | secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty.<br/>More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | false |
| user | string | user is optional: User is the rados user name, default is admin<br/>More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].cephfs.secretRef

secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty.<br/>More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].cinder

cinder represents a cinder volume attached and mounted on kubelets host machine.<br/>Deprecated: Cinder is deprecated. All operations for the in-tree cinder type<br/>are redirected to the cinder.csi.openstack.org CSI driver.<br/>More info: https://examples.k8s.io/mysql-cinder-pd/README.md

| Name | Type | Description | Required |
|------|------|-------------|----------|
| volumeID | string | volumeID used to identify the volume in cinder.<br/>More info: https://examples.k8s.io/mysql-cinder-pd/README.md | true |
| fsType | string | fsType is the filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br/>More info: https://examples.k8s.io/mysql-cinder-pd/README.md | false |
| readOnly | boolean | readOnly defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts.<br/>More info: https://examples.k8s.io/mysql-cinder-pd/README.md | false |
| [secretRef](#lmdeploymentspecopenwebuipipelinesvolumesindexcindersecretref) | object | secretRef is optional: points to a secret object containing parameters used to connect<br/>to OpenStack. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].cinder.secretRef

secretRef is optional: points to a secret object containing parameters used to connect<br/>to OpenStack.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].configMap

configMap represents a configMap that should populate this volume

| Name | Type | Description | Required |
|------|------|-------------|----------|
| defaultMode | integer | defaultMode is optional: mode bits used to set permissions on created files by default.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>Defaults to 0644.<br/>Directories within the path are not affected by this setting.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
| [items](#lmdeploymentspecopenwebuipipelinesvolumesindexconfigmapitemsindex) | []object | items if unspecified, each key-value pair in the Data field of the referenced<br/>ConfigMap will be projected into the volume as a file whose name is the<br/>key and content is the value. If specified, the listed keys will be<br/>projected into the specified paths, and unlisted keys will not be<br/>present. If a key is specified which is not present in the ConfigMap,<br/>the volume setup will error unless it is marked optional. Paths must be<br/>relative and may not contain the '..' path or start with '..'. | false |
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
| optional | boolean | optional specify whether the ConfigMap or its keys must be defined | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].configMap.items[index]

Maps a string key to a path within a volume.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the key to project. | true |
| path | string | path is the relative path of the file to map the key to.<br/>May not be an absolute path.<br/>May not contain the path element '..'.<br/>May not start with the string '..'. | true |
| mode | integer | mode is Optional: mode bits used to set permissions on this file.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>If not specified, the volume defaultMode will be used.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].csi

csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| driver | string | driver is the name of the CSI driver that handles this volume.<br/>Consult with your admin for the correct name as registered in the cluster. | true |
| fsType | string | fsType to mount. Ex. "ext4", "xfs", "ntfs".<br/>If not provided, the empty value is passed to the associated CSI driver<br/>which will determine the default filesystem to apply. | false |
| [nodePublishSecretRef](#lmdeploymentspecopenwebuipipelinesvolumesindexcsinodepublishsecretref) | object | nodePublishSecretRef is a reference to the secret object containing<br/>sensitive information to pass to the CSI driver to complete the CSI<br/>NodePublishVolume and NodeUnpublishVolume calls.<br/>This field is optional, and  may be empty if no secret is required. If the<br/>secret object contains more than one secret, all secret references are passed. | false |
| readOnly | boolean | readOnly specifies a read-only configuration for the volume.<br/>Defaults to false (read/write). | false |
| volumeAttributes | map[string]string | volumeAttributes stores driver-specific properties that are passed to the CSI<br/>driver. Consult your driver's documentation for supported values. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].csi.nodePublishSecretRef

nodePublishSecretRef is a reference to the secret object containing<br/>sensitive information to pass to the CSI driver to complete the CSI<br/>NodePublishVolume and NodeUnpublishVolume calls.<br/>This field is optional, and  may be empty if no secret is required. If the<br/>secret object contains more than one secret, all secret references are passed.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].downwardAPI

downwardAPI represents downward API about the pod that should populate this volume

| Name | Type | Description | Required |
|------|------|-------------|----------|
| defaultMode | integer | Optional: mode bits to use on created files by default. Must be a<br/>Optional: mode bits used to set permissions on created files by default.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>Defaults to 0644.<br/>Directories within the path are not affected by this setting.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
| [items](#lmdeploymentspecopenwebuipipelinesvolumesindexdownwardapiitemsindex) | []object | Items is a list of downward API volume file | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].downwardAPI.items[index]

DownwardAPIVolumeFile represents information to create the file containing the pod field

| Name | Type | Description | Required |
|------|------|-------------|----------|
| path | string | Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..' | true |
| [fieldRef](#lmdeploymentspecopenwebuipipelinesvolumesindexdownwardapiitemsindexfieldref) | object | Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported. | false |
| mode | integer | Optional: mode bits used to set permissions on this file, must be an octal value<br/>between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>If not specified, the volume defaultMode will be used.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
| [resourceFieldRef](#lmdeploymentspecopenwebuipipelinesvolumesindexdownwardapiitemsindexresourcefieldref) | object | Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].downwardAPI.items[index].fieldRef

Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| fieldPath | string | Path of the field to select in the specified API version. | true |
| apiVersion | string | Version of the schema the FieldPath is written in terms of, defaults to "v1". | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].downwardAPI.items[index].resourceFieldRef

Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| resource | string | Required: resource to select | true |
| containerName | string | Container name: required for volumes, optional for env vars | false |
| divisor | int or string | Specifies the output format of the exposed resources, defaults to "1" | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].emptyDir

emptyDir represents a temporary directory that shares a pod's lifetime.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir

| Name | Type | Description | Required |
|------|------|-------------|----------|
| medium | string | medium represents what type of storage medium should back this directory.<br/>The default is "" which means to use the node's default medium.<br/>Must be an empty string (default) or Memory.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir | false |
| sizeLimit | int or string | sizeLimit is the total amount of local storage required for this EmptyDir volume.<br/>The size limit is also applicable for memory medium.<br/>The maximum usage on memory medium EmptyDir would be the minimum value between<br/>the SizeLimit specified here and the sum of memory limits of all containers in a pod.<br/>The default is nil which means that the limit is undefined.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].ephemeral

ephemeral represents a volume that is handled by a cluster storage driver.<br/>The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts,<br/>and deleted when the pod is removed.<br/><br/>Use this if:<br/>a) the volume is only needed while the pod runs,<br/>b) features of normal volumes like restoring from snapshot or capacity<br/>   tracking are needed,<br/>c) the storage driver is specified through a storage class, and<br/>d) the storage driver supports dynamic volume provisioning through<br/>   a PersistentVolumeClaim (see EphemeralVolumeSource for more<br/>   information on the connection between this volume type<br/>   and PersistentVolumeClaim).<br/><br/>Use PersistentVolumeClaim or one of the vendor-specific<br/>APIs for volumes that persist for longer than the lifecycle<br/>of an individual pod.<br/><br/>Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to<br/>be used that way - see the documentation of the driver for<br/>more information.<br/><br/>A pod can use both types of ephemeral volumes and<br/>persistent volumes at the same time.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [volumeClaimTemplate](#lmdeploymentspecopenwebuipipelinesvolumesindexephemeralvolumeclaimtemplate) | object | Will be used to create a stand-alone PVC to provision the volume.<br/>The pod in which this EphemeralVolumeSource is embedded will be the<br/>owner of the PVC, i.e. the PVC will be deleted together with the<br/>pod.  The name of the PVC will be `&lt;pod name&gt;-&lt;volume name&gt;` where<br/>`&lt;volume name&gt;` is the name from the `PodSpec.Volumes` array<br/>entry. Pod validation will reject the pod if the concatenated name<br/>is not valid for a PVC (for example, too long).<br/><br/>An existing PVC with that name that is not owned by the pod<br/>will *not* be used for the pod to avoid using an unrelated<br/>volume by mistake. Starting the pod is then blocked until<br/>the unrelated PVC is removed. If such a pre-created PVC is<br/>meant to be used by the pod, the PVC has to updated with an<br/>owner reference to the pod once the pod exists. Normally<br/>this should not be necessary, but it may be useful when<br/>manually reconstructing a broken cluster.<br/><br/>This field is read-only and no changes will be made by Kubernetes<br/>to the PVC after it has been created.<br/><br/>Required, must not be nil. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].ephemeral.volumeClaimTemplate

Will be used to create a stand-alone PVC to provision the volume.<br/>The pod in which this EphemeralVolumeSource is embedded will be the<br/>owner of the PVC, i.e. the PVC will be deleted together with the<br/>pod.  The name of the PVC will be `&lt;pod name&gt;-&lt;volume name&gt;` where<br/>`&lt;volume name&gt;` is the name from the `PodSpec.Volumes` array<br/>entry. Pod validation will reject the pod if the concatenated name<br/>is not valid for a PVC (for example, too long).<br/><br/>An existing PVC with that name that is not owned by the pod<br/>will *not* be used for the pod to avoid using an unrelated<br/>volume by mistake. Starting the pod is then blocked until<br/>the unrelated PVC is removed. If such a pre-created PVC is<br/>meant to be used by the pod, the PVC has to updated with an<br/>owner reference to the pod once the pod exists. Normally<br/>this should not be necessary, but it may be useful when<br/>manually reconstructing a broken cluster.<br/><br/>This field is read-only and no changes will be made by Kubernetes<br/>to the PVC after it has been created.<br/><br/>Required, must not be nil.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [spec](#lmdeploymentspecopenwebuipipelinesvolumesindexephemeralvolumeclaimtemplatespec) | object | The specification for the PersistentVolumeClaim. The entire content is<br/>copied unchanged into the PVC that gets created from this<br/>template. The same fields as in a PersistentVolumeClaim<br/>are also valid here. | true |
| metadata | object | May contain labels and annotations that will be copied into the PVC<br/>when creating it. No other fields are allowed and will be rejected during<br/>validation. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].ephemeral.volumeClaimTemplate.spec

The specification for the PersistentVolumeClaim. The entire content is<br/>copied unchanged into the PVC that gets created from this<br/>template. The same fields as in a PersistentVolumeClaim<br/>are also valid here.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| accessModes | []string | accessModes contains the desired access modes the volume should have.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 | false |
| [dataSource](#lmdeploymentspecopenwebuipipelinesvolumesindexephemeralvolumeclaimtemplatespecdatasource) | object | dataSource field can be used to specify either:<br/>* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)<br/>* An existing PVC (PersistentVolumeClaim)<br/>If the provisioner or an external controller can support the specified data source,<br/>it will create a new volume based on the contents of the specified data source.<br/>When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,<br/>and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.<br/>If the namespace is specified, then dataSourceRef will not be copied to dataSource. | false |
| [dataSourceRef](#lmdeploymentspecopenwebuipipelinesvolumesindexephemeralvolumeclaimtemplatespecdatasourceref) | object | dataSourceRef specifies the object from which to populate the volume with data, if a non-empty<br/>volume is desired. This may be any object from a non-empty API group (non<br/>core object) or a PersistentVolumeClaim object.<br/>When this field is specified, volume binding will only succeed if the type of<br/>the specified object matches some installed volume populator or dynamic<br/>provisioner.<br/>This field will replace the functionality of the dataSource field and as such<br/>if both fields are non-empty, they must have the same value. For backwards<br/>compatibility, when namespace isn't specified in dataSourceRef,<br/>both fields (dataSource and dataSourceRef) will be set to the same<br/>value automatically if one of them is empty and the other is non-empty.<br/>When namespace is specified in dataSourceRef,<br/>dataSource isn't set to the same value and must be empty.<br/>There are three important differences between dataSource and dataSourceRef:<br/>* While dataSource only allows two specific types of objects, dataSourceRef<br/>  allows any non-core object, as well as PersistentVolumeClaim objects.<br/>* While dataSource ignores disallowed values (dropping them), dataSourceRef<br/>  preserves all values, and generates an error if a disallowed value is<br/>  specified.<br/>* While dataSource only allows local objects, dataSourceRef allows objects<br/>  in any namespaces.<br/>(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.<br/>(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled. | false |
| [resources](#lmdeploymentspecopenwebuipipelinesvolumesindexephemeralvolumeclaimtemplatespecresources) | object | resources represents the minimum resources the volume should have.<br/>If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements<br/>that are lower than previous value but must still be higher than capacity recorded in the<br/>status field of the claim.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources | false |
| [selector](#lmdeploymentspecopenwebuipipelinesvolumesindexephemeralvolumeclaimtemplatespecselector) | object | selector is a label query over volumes to consider for binding. | false |
| storageClassName | string | storageClassName is the name of the StorageClass required by the claim.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1 | false |
| volumeAttributesClassName | string | volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.<br/>If specified, the CSI driver will create or update the volume with the attributes defined<br/>in the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,<br/>it can be changed after the claim is created. An empty string value means that no VolumeAttributesClass<br/>will be applied to the claim but it's not allowed to reset this field to empty string once it is set.<br/>If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClass<br/>will be set by the persistentvolume controller if it exists.<br/>If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will be<br/>set to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resource<br/>exists.<br/>More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/<br/>(Beta) Using this field requires the VolumeAttributesClass feature gate to be enabled (off by default). | false |
| volumeMode | string | volumeMode defines what type of volume is required by the claim.<br/>Value of Filesystem is implied when not included in claim spec. | false |
| volumeName | string | volumeName is the binding reference to the PersistentVolume backing this claim. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].ephemeral.volumeClaimTemplate.spec.dataSource

dataSource field can be used to specify either:<br/>* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)<br/>* An existing PVC (PersistentVolumeClaim)<br/>If the provisioner or an external controller can support the specified data source,<br/>it will create a new volume based on the contents of the specified data source.<br/>When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,<br/>and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.<br/>If the namespace is specified, then dataSourceRef will not be copied to dataSource.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| kind | string | Kind is the type of resource being referenced | true |
| name | string | Name is the name of resource being referenced | true |
| apiGroup | string | APIGroup is the group for the resource being referenced.<br/>If APIGroup is not specified, the specified Kind must be in the core API group.<br/>For any other third-party types, APIGroup is required. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].ephemeral.volumeClaimTemplate.spec.dataSourceRef

dataSourceRef specifies the object from which to populate the volume with data, if a non-empty<br/>volume is desired. This may be any object from a non-empty API group (non<br/>core object) or a PersistentVolumeClaim object.<br/>When this field is specified, volume binding will only succeed if the type of<br/>the specified object matches some installed volume populator or dynamic<br/>provisioner.<br/>This field will replace the functionality of the dataSource field and as such<br/>if both fields are non-empty, they must have the same value. For backwards<br/>compatibility, when namespace isn't specified in dataSourceRef,<br/>both fields (dataSource and dataSourceRef) will be set to the same<br/>value automatically if one of them is empty and the other is non-empty.<br/>When namespace is specified in dataSourceRef,<br/>dataSource isn't set to the same value and must be empty.<br/>There are three important differences between dataSource and dataSourceRef:<br/>* While dataSource only allows two specific types of objects, dataSourceRef<br/>  allows any non-core object, as well as PersistentVolumeClaim objects.<br/>* While dataSource ignores disallowed values (dropping them), dataSourceRef<br/>  preserves all values, and generates an error if a disallowed value is<br/>  specified.<br/>* While dataSource only allows local objects, dataSourceRef allows objects<br/>  in any namespaces.<br/>(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.<br/>(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| kind | string | Kind is the type of resource being referenced | true |
| name | string | Name is the name of resource being referenced | true |
| apiGroup | string | APIGroup is the group for the resource being referenced.<br/>If APIGroup is not specified, the specified Kind must be in the core API group.<br/>For any other third-party types, APIGroup is required. | false |
| namespace | string | Namespace is the namespace of resource being referenced<br/>Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.<br/>(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].ephemeral.volumeClaimTemplate.spec.resources

resources represents the minimum resources the volume should have.<br/>If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements<br/>that are lower than previous value but must still be higher than capacity recorded in the<br/>status field of the claim.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources

| Name | Type | Description | Required |
|------|------|-------------|----------|
| limits | map[string]int or string | Limits describes the maximum amount of compute resources allowed.<br/>More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ | false |
| requests | map[string]int or string | Requests describes the minimum amount of compute resources required.<br/>If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,<br/>otherwise to an implementation-defined value. Requests cannot exceed Limits.<br/>More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].ephemeral.volumeClaimTemplate.spec.selector

selector is a label query over volumes to consider for binding.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecopenwebuipipelinesvolumesindexephemeralvolumeclaimtemplatespecselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].ephemeral.volumeClaimTemplate.spec.selector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].fc

fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| fsType | string | fsType is the filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. | false |
| lun | integer | lun is Optional: FC target lun number<br/>*Format*: int32<br/> | false |
| readOnly | boolean | readOnly is Optional: Defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts. | false |
| targetWWNs | []string | targetWWNs is Optional: FC target worldwide names (WWNs) | false |
| wwids | []string | wwids Optional: FC volume world wide identifiers (wwids)<br/>Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].flexVolume

flexVolume represents a generic volume resource that is<br/>provisioned/attached using an exec based plugin.<br/>Deprecated: FlexVolume is deprecated. Consider using a CSIDriver instead.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| driver | string | driver is the name of the driver to use for this volume. | true |
| fsType | string | fsType is the filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script. | false |
| options | map[string]string | options is Optional: this field holds extra command options if any. | false |
| readOnly | boolean | readOnly is Optional: defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts. | false |
| [secretRef](#lmdeploymentspecopenwebuipipelinesvolumesindexflexvolumesecretref) | object | secretRef is Optional: secretRef is reference to the secret object containing<br/>sensitive information to pass to the plugin scripts. This may be<br/>empty if no secret object is specified. If the secret object<br/>contains more than one secret, all secrets are passed to the plugin<br/>scripts. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].flexVolume.secretRef

secretRef is Optional: secretRef is reference to the secret object containing<br/>sensitive information to pass to the plugin scripts. This may be<br/>empty if no secret object is specified. If the secret object<br/>contains more than one secret, all secrets are passed to the plugin<br/>scripts.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].flocker

flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running.<br/>Deprecated: Flocker is deprecated and the in-tree flocker type is no longer supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| datasetName | string | datasetName is Name of the dataset stored as metadata -&gt; name on the dataset for Flocker<br/>should be considered as deprecated | false |
| datasetUUID | string | datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].gcePersistentDisk

gcePersistentDisk represents a GCE Disk resource that is attached to a<br/>kubelet's host machine and then exposed to the pod.<br/>Deprecated: GCEPersistentDisk is deprecated. All operations for the in-tree<br/>gcePersistentDisk type are redirected to the pd.csi.storage.gke.io CSI driver.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk

| Name | Type | Description | Required |
|------|------|-------------|----------|
| pdName | string | pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk | true |
| fsType | string | fsType is filesystem type of the volume that you want to mount.<br/>Tip: Ensure that the filesystem type is supported by the host operating system.<br/>Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk | false |
| partition | integer | partition is the partition in the volume that you want to mount.<br/>If omitted, the default is to mount by volume name.<br/>Examples: For volume /dev/sda1, you specify the partition as "1".<br/>Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk<br/>*Format*: int32<br/> | false |
| readOnly | boolean | readOnly here will force the ReadOnly setting in VolumeMounts.<br/>Defaults to false.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].gitRepo

gitRepo represents a git repository at a particular revision.<br/>Deprecated: GitRepo is deprecated. To provision a container with a git repo, mount an<br/>EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir<br/>into the Pod's container.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| repository | string | repository is the URL | true |
| directory | string | directory is the target directory name.<br/>Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the<br/>git repository.  Otherwise, if specified, the volume will contain the git repository in<br/>the subdirectory with the given name. | false |
| revision | string | revision is the commit hash for the specified revision. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].glusterfs

glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.<br/>Deprecated: Glusterfs is deprecated and the in-tree glusterfs type is no longer supported.<br/>More info: https://examples.k8s.io/volumes/glusterfs/README.md

| Name | Type | Description | Required |
|------|------|-------------|----------|
| endpoints | string | endpoints is the endpoint name that details Glusterfs topology.<br/>More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod | true |
| path | string | path is the Glusterfs volume path.<br/>More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod | true |
| readOnly | boolean | readOnly here will force the Glusterfs volume to be mounted with read-only permissions.<br/>Defaults to false.<br/>More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].hostPath

hostPath represents a pre-existing file or directory on the host<br/>machine that is directly exposed to the container. This is generally<br/>used for system agents or other privileged things that are allowed<br/>to see the host machine. Most containers will NOT need this.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath

| Name | Type | Description | Required |
|------|------|-------------|----------|
| path | string | path of the directory on the host.<br/>If the path is a symlink, it will follow the link to the real path.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath | true |
| type | string | type for HostPath Volume<br/>Defaults to ""<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].image

image represents an OCI object (a container image or artifact) pulled and mounted on the kubelet's host machine.<br/>The volume is resolved at pod startup depending on which PullPolicy value is provided:<br/><br/>- Always: the kubelet always attempts to pull the reference. Container creation will fail If the pull fails.<br/>- Never: the kubelet never pulls the reference and only uses a local image or artifact. Container creation will fail if the reference isn't present.<br/>- IfNotPresent: the kubelet pulls if the reference isn't already present on disk. Container creation will fail if the reference isn't present and the pull fails.<br/><br/>The volume gets re-resolved if the pod gets deleted and recreated, which means that new remote content will become available on pod recreation.<br/>A failure to resolve or pull the image during pod startup will block containers from starting and may add significant latency. Failures will be retried using normal volume backoff and will be reported on the pod reason and message.<br/>The types of objects that may be mounted by this volume are defined by the container runtime implementation on a host machine and at minimum must include all valid types supported by the container image field.<br/>The OCI object gets mounted in a single directory (spec.containers[*].volumeMounts.mountPath) by merging the manifest layers in the same way as for container images.<br/>The volume will be mounted read-only (ro) and non-executable files (noexec).<br/>Sub path mounts for containers are not supported (spec.containers[*].volumeMounts.subpath) before 1.33.<br/>The field spec.securityContext.fsGroupChangePolicy has no effect on this volume type.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| pullPolicy | string | Policy for pulling OCI objects. Possible values are:<br/>Always: the kubelet always attempts to pull the reference. Container creation will fail If the pull fails.<br/>Never: the kubelet never pulls the reference and only uses a local image or artifact. Container creation will fail if the reference isn't present.<br/>IfNotPresent: the kubelet pulls if the reference isn't already present on disk. Container creation will fail if the reference isn't present and the pull fails.<br/>Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. | false |
| reference | string | Required: Image or artifact reference to be used.<br/>Behaves in the same way as pod.spec.containers[*].image.<br/>Pull secrets will be assembled in the same way as for the container image by looking up node credentials, SA image pull secrets, and pod spec image pull secrets.<br/>More info: https://kubernetes.io/docs/concepts/containers/images<br/>This field is optional to allow higher level config management to default or override<br/>container images in workload controllers like Deployments and StatefulSets. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].iscsi

iscsi represents an ISCSI Disk resource that is attached to a<br/>kubelet's host machine and then exposed to the pod.<br/>More info: https://examples.k8s.io/volumes/iscsi/README.md

| Name | Type | Description | Required |
|------|------|-------------|----------|
| iqn | string | iqn is the target iSCSI Qualified Name. | true |
| lun | integer | lun represents iSCSI Target Lun number.<br/>*Format*: int32<br/> | true |
| targetPortal | string | targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port<br/>is other than default (typically TCP ports 860 and 3260). | true |
| chapAuthDiscovery | boolean | chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication | false |
| chapAuthSession | boolean | chapAuthSession defines whether support iSCSI Session CHAP authentication | false |
| fsType | string | fsType is the filesystem type of the volume that you want to mount.<br/>Tip: Ensure that the filesystem type is supported by the host operating system.<br/>Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi | false |
| initiatorName | string | initiatorName is the custom iSCSI Initiator Name.<br/>If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface<br/>&lt;target portal&gt;:&lt;volume name&gt; will be created for the connection. | false |
| iscsiInterface | string | iscsiInterface is the interface Name that uses an iSCSI transport.<br/>Defaults to 'default' (tcp).<br/>*Default*: default<br/> | false |
| portals | []string | portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port<br/>is other than default (typically TCP ports 860 and 3260). | false |
| readOnly | boolean | readOnly here will force the ReadOnly setting in VolumeMounts.<br/>Defaults to false. | false |
| [secretRef](#lmdeploymentspecopenwebuipipelinesvolumesindexiscsisecretref) | object | secretRef is the CHAP Secret for iSCSI target and initiator authentication | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].iscsi.secretRef

secretRef is the CHAP Secret for iSCSI target and initiator authentication

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].nfs

nfs represents an NFS mount on the host that shares a pod's lifetime<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs

| Name | Type | Description | Required |
|------|------|-------------|----------|
| path | string | path that is exported by the NFS server.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs | true |
| server | string | server is the hostname or IP address of the NFS server.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs | true |
| readOnly | boolean | readOnly here will force the NFS export to be mounted with read-only permissions.<br/>Defaults to false.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].persistentVolumeClaim

persistentVolumeClaimVolumeSource represents a reference to a<br/>PersistentVolumeClaim in the same namespace.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims

| Name | Type | Description | Required |
|------|------|-------------|----------|
| claimName | string | claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims | true |
| readOnly | boolean | readOnly Will force the ReadOnly setting in VolumeMounts.<br/>Default false. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].photonPersistentDisk

photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine.<br/>Deprecated: PhotonPersistentDisk is deprecated and the in-tree photonPersistentDisk type is no longer supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| pdID | string | pdID is the ID that identifies Photon Controller persistent disk | true |
| fsType | string | fsType is the filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].portworxVolume

portworxVolume represents a portworx volume attached and mounted on kubelets host machine.<br/>Deprecated: PortworxVolume is deprecated. All operations for the in-tree portworxVolume type<br/>are redirected to the pxd.portworx.com CSI driver when the CSIMigrationPortworx feature-gate<br/>is on.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| volumeID | string | volumeID uniquely identifies a Portworx volume | true |
| fsType | string | fSType represents the filesystem type to mount<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified. | false |
| readOnly | boolean | readOnly defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].projected

projected items for all in one resources secrets, configmaps, and downward API

| Name | Type | Description | Required |
|------|------|-------------|----------|
| defaultMode | integer | defaultMode are the mode bits used to set permissions on created files by default.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>Directories within the path are not affected by this setting.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
| [sources](#lmdeploymentspecopenwebuipipelinesvolumesindexprojectedsourcesindex) | []object | sources is the list of volume projections. Each entry in this list<br/>handles one source. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].projected.sources[index]

Projection that may be projected along with other supported volume types.<br/>Exactly one of these fields must be set.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [clusterTrustBundle](#lmdeploymentspecopenwebuipipelinesvolumesindexprojectedsourcesindexclustertrustbundle) | object | ClusterTrustBundle allows a pod to access the `.spec.trustBundle` field<br/>of ClusterTrustBundle objects in an auto-updating file.<br/><br/>Alpha, gated by the ClusterTrustBundleProjection feature gate.<br/><br/>ClusterTrustBundle objects can either be selected by name, or by the<br/>combination of signer name and a label selector.<br/><br/>Kubelet performs aggressive normalization of the PEM contents written<br/>into the pod filesystem.  Esoteric PEM features such as inter-block<br/>comments and block headers are stripped.  Certificates are deduplicated.<br/>The ordering of certificates within the file is arbitrary, and Kubelet<br/>may change the order over time. | false |
| [configMap](#lmdeploymentspecopenwebuipipelinesvolumesindexprojectedsourcesindexconfigmap) | object | configMap information about the configMap data to project | false |
| [downwardAPI](#lmdeploymentspecopenwebuipipelinesvolumesindexprojectedsourcesindexdownwardapi) | object | downwardAPI information about the downwardAPI data to project | false |
| [secret](#lmdeploymentspecopenwebuipipelinesvolumesindexprojectedsourcesindexsecret) | object | secret information about the secret data to project | false |
| [serviceAccountToken](#lmdeploymentspecopenwebuipipelinesvolumesindexprojectedsourcesindexserviceaccounttoken) | object | serviceAccountToken is information about the serviceAccountToken data to project | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].projected.sources[index].clusterTrustBundle

ClusterTrustBundle allows a pod to access the `.spec.trustBundle` field<br/>of ClusterTrustBundle objects in an auto-updating file.<br/><br/>Alpha, gated by the ClusterTrustBundleProjection feature gate.<br/><br/>ClusterTrustBundle objects can either be selected by name, or by the<br/>combination of signer name and a label selector.<br/><br/>Kubelet performs aggressive normalization of the PEM contents written<br/>into the pod filesystem.  Esoteric PEM features such as inter-block<br/>comments and block headers are stripped.  Certificates are deduplicated.<br/>The ordering of certificates within the file is arbitrary, and Kubelet<br/>may change the order over time.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| path | string | Relative path from the volume root to write the bundle. | true |
| [labelSelector](#lmdeploymentspecopenwebuipipelinesvolumesindexprojectedsourcesindexclustertrustbundlelabelselector) | object | Select all ClusterTrustBundles that match this label selector.  Only has<br/>effect if signerName is set.  Mutually-exclusive with name.  If unset,<br/>interpreted as "match nothing".  If set but empty, interpreted as "match<br/>everything". | false |
| name | string | Select a single ClusterTrustBundle by object name.  Mutually-exclusive<br/>with signerName and labelSelector. | false |
| optional | boolean | If true, don't block pod startup if the referenced ClusterTrustBundle(s)<br/>aren't available.  If using name, then the named ClusterTrustBundle is<br/>allowed not to exist.  If using signerName, then the combination of<br/>signerName and labelSelector is allowed to match zero<br/>ClusterTrustBundles. | false |
| signerName | string | Select all ClusterTrustBundles that match this signer name.<br/>Mutually-exclusive with name.  The contents of all selected<br/>ClusterTrustBundles will be unified and deduplicated. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].projected.sources[index].clusterTrustBundle.labelSelector

Select all ClusterTrustBundles that match this label selector.  Only has<br/>effect if signerName is set.  Mutually-exclusive with name.  If unset,<br/>interpreted as "match nothing".  If set but empty, interpreted as "match<br/>everything".

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecopenwebuipipelinesvolumesindexprojectedsourcesindexclustertrustbundlelabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].projected.sources[index].clusterTrustBundle.labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].projected.sources[index].configMap

configMap information about the configMap data to project

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [items](#lmdeploymentspecopenwebuipipelinesvolumesindexprojectedsourcesindexconfigmapitemsindex) | []object | items if unspecified, each key-value pair in the Data field of the referenced<br/>ConfigMap will be projected into the volume as a file whose name is the<br/>key and content is the value. If specified, the listed keys will be<br/>projected into the specified paths, and unlisted keys will not be<br/>present. If a key is specified which is not present in the ConfigMap,<br/>the volume setup will error unless it is marked optional. Paths must be<br/>relative and may not contain the '..' path or start with '..'. | false |
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
| optional | boolean | optional specify whether the ConfigMap or its keys must be defined | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].projected.sources[index].configMap.items[index]

Maps a string key to a path within a volume.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the key to project. | true |
| path | string | path is the relative path of the file to map the key to.<br/>May not be an absolute path.<br/>May not contain the path element '..'.<br/>May not start with the string '..'. | true |
| mode | integer | mode is Optional: mode bits used to set permissions on this file.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>If not specified, the volume defaultMode will be used.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].projected.sources[index].downwardAPI

downwardAPI information about the downwardAPI data to project

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [items](#lmdeploymentspecopenwebuipipelinesvolumesindexprojectedsourcesindexdownwardapiitemsindex) | []object | Items is a list of DownwardAPIVolume file | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].projected.sources[index].downwardAPI.items[index]

DownwardAPIVolumeFile represents information to create the file containing the pod field

| Name | Type | Description | Required |
|------|------|-------------|----------|
| path | string | Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..' | true |
| [fieldRef](#lmdeploymentspecopenwebuipipelinesvolumesindexprojectedsourcesindexdownwardapiitemsindexfieldref) | object | Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported. | false |
| mode | integer | Optional: mode bits used to set permissions on this file, must be an octal value<br/>between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>If not specified, the volume defaultMode will be used.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
| [resourceFieldRef](#lmdeploymentspecopenwebuipipelinesvolumesindexprojectedsourcesindexdownwardapiitemsindexresourcefieldref) | object | Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].projected.sources[index].downwardAPI.items[index].fieldRef

Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| fieldPath | string | Path of the field to select in the specified API version. | true |
| apiVersion | string | Version of the schema the FieldPath is written in terms of, defaults to "v1". | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].projected.sources[index].downwardAPI.items[index].resourceFieldRef

Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| resource | string | Required: resource to select | true |
| containerName | string | Container name: required for volumes, optional for env vars | false |
| divisor | int or string | Specifies the output format of the exposed resources, defaults to "1" | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].projected.sources[index].secret

secret information about the secret data to project

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [items](#lmdeploymentspecopenwebuipipelinesvolumesindexprojectedsourcesindexsecretitemsindex) | []object | items if unspecified, each key-value pair in the Data field of the referenced<br/>Secret will be projected into the volume as a file whose name is the<br/>key and content is the value. If specified, the listed keys will be<br/>projected into the specified paths, and unlisted keys will not be<br/>present. If a key is specified which is not present in the Secret,<br/>the volume setup will error unless it is marked optional. Paths must be<br/>relative and may not contain the '..' path or start with '..'. | false |
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
| optional | boolean | optional field specify whether the Secret or its key must be defined | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].projected.sources[index].secret.items[index]

Maps a string key to a path within a volume.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the key to project. | true |
| path | string | path is the relative path of the file to map the key to.<br/>May not be an absolute path.<br/>May not contain the path element '..'.<br/>May not start with the string '..'. | true |
| mode | integer | mode is Optional: mode bits used to set permissions on this file.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>If not specified, the volume defaultMode will be used.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].projected.sources[index].serviceAccountToken

serviceAccountToken is information about the serviceAccountToken data to project

| Name | Type | Description | Required |
|------|------|-------------|----------|
| path | string | path is the path relative to the mount point of the file to project the<br/>token into. | true |
| audience | string | audience is the intended audience of the token. A recipient of a token<br/>must identify itself with an identifier specified in the audience of the<br/>token, and otherwise should reject the token. The audience defaults to the<br/>identifier of the apiserver. | false |
| expirationSeconds | integer | expirationSeconds is the requested duration of validity of the service<br/>account token. As the token approaches expiration, the kubelet volume<br/>plugin will proactively rotate the service account token. The kubelet will<br/>start trying to rotate the token if the token is older than 80 percent of<br/>its time to live or if the token is older than 24 hours.Defaults to 1 hour<br/>and must be at least 10 minutes.<br/>*Format*: int64<br/> | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].quobyte

quobyte represents a Quobyte mount on the host that shares a pod's lifetime.<br/>Deprecated: Quobyte is deprecated and the in-tree quobyte type is no longer supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| registry | string | registry represents a single or multiple Quobyte Registry services<br/>specified as a string as host:port pair (multiple entries are separated with commas)<br/>which acts as the central registry for volumes | true |
| volume | string | volume is a string that references an already created Quobyte volume by name. | true |
| group | string | group to map volume access to<br/>Default is no group | false |
| readOnly | boolean | readOnly here will force the Quobyte volume to be mounted with read-only permissions.<br/>Defaults to false. | false |
| tenant | string | tenant owning the given Quobyte volume in the Backend<br/>Used with dynamically provisioned Quobyte volumes, value is set by the plugin | false |
| user | string | user to map volume access to<br/>Defaults to serivceaccount user | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].rbd

rbd represents a Rados Block Device mount on the host that shares a pod's lifetime.<br/>Deprecated: RBD is deprecated and the in-tree rbd type is no longer supported.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md

| Name | Type | Description | Required |
|------|------|-------------|----------|
| image | string | image is the rados image name.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | true |
| monitors | []string | monitors is a collection of Ceph monitors.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | true |
| fsType | string | fsType is the filesystem type of the volume that you want to mount.<br/>Tip: Ensure that the filesystem type is supported by the host operating system.<br/>Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd | false |
| keyring | string | keyring is the path to key ring for RBDUser.<br/>Default is /etc/ceph/keyring.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it<br/>*Default*: /etc/ceph/keyring<br/> | false |
| pool | string | pool is the rados pool name.<br/>Default is rbd.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it<br/>*Default*: rbd<br/> | false |
| readOnly | boolean | readOnly here will force the ReadOnly setting in VolumeMounts.<br/>Defaults to false.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | false |
| [secretRef](#lmdeploymentspecopenwebuipipelinesvolumesindexrbdsecretref) | object | secretRef is name of the authentication secret for RBDUser. If provided<br/>overrides keyring.<br/>Default is nil.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | false |
| user | string | user is the rados user name.<br/>Default is admin.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it<br/>*Default*: admin<br/> | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].rbd.secretRef

secretRef is name of the authentication secret for RBDUser. If provided<br/>overrides keyring.<br/>Default is nil.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].scaleIO

scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.<br/>Deprecated: ScaleIO is deprecated and the in-tree scaleIO type is no longer supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| gateway | string | gateway is the host address of the ScaleIO API Gateway. | true |
| [secretRef](#lmdeploymentspecopenwebuipipelinesvolumesindexscaleiosecretref) | object | secretRef references to the secret for ScaleIO user and other<br/>sensitive information. If this is not provided, Login operation will fail. | true |
| system | string | system is the name of the storage system as configured in ScaleIO. | true |
| fsType | string | fsType is the filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs".<br/>Default is "xfs".<br/>*Default*: xfs<br/> | false |
| protectionDomain | string | protectionDomain is the name of the ScaleIO Protection Domain for the configured storage. | false |
| readOnly | boolean | readOnly Defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts. | false |
| sslEnabled | boolean | sslEnabled Flag enable/disable SSL communication with Gateway, default false | false |
| storageMode | string | storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.<br/>Default is ThinProvisioned.<br/>*Default*: ThinProvisioned<br/> | false |
| storagePool | string | storagePool is the ScaleIO Storage Pool associated with the protection domain. | false |
| volumeName | string | volumeName is the name of a volume already created in the ScaleIO system<br/>that is associated with this volume source. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].scaleIO.secretRef

secretRef references to the secret for ScaleIO user and other<br/>sensitive information. If this is not provided, Login operation will fail.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].secret

secret represents a secret that should populate this volume.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#secret

| Name | Type | Description | Required |
|------|------|-------------|----------|
| defaultMode | integer | defaultMode is Optional: mode bits used to set permissions on created files by default.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values<br/>for mode bits. Defaults to 0644.<br/>Directories within the path are not affected by this setting.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
| [items](#lmdeploymentspecopenwebuipipelinesvolumesindexsecretitemsindex) | []object | items If unspecified, each key-value pair in the Data field of the referenced<br/>Secret will be projected into the volume as a file whose name is the<br/>key and content is the value. If specified, the listed keys will be<br/>projected into the specified paths, and unlisted keys will not be<br/>present. If a key is specified which is not present in the Secret,<br/>the volume setup will error unless it is marked optional. Paths must be<br/>relative and may not contain the '..' path or start with '..'. | false |
| optional | boolean | optional field specify whether the Secret or its keys must be defined | false |
| secretName | string | secretName is the name of the secret in the pod's namespace to use.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#secret | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].secret.items[index]

Maps a string key to a path within a volume.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the key to project. | true |
| path | string | path is the relative path of the file to map the key to.<br/>May not be an absolute path.<br/>May not contain the path element '..'.<br/>May not start with the string '..'. | true |
| mode | integer | mode is Optional: mode bits used to set permissions on this file.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>If not specified, the volume defaultMode will be used.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].storageos

storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.<br/>Deprecated: StorageOS is deprecated and the in-tree storageos type is no longer supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| fsType | string | fsType is the filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. | false |
| readOnly | boolean | readOnly defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts. | false |
| [secretRef](#lmdeploymentspecopenwebuipipelinesvolumesindexstorageossecretref) | object | secretRef specifies the secret to use for obtaining the StorageOS API<br/>credentials.  If not specified, default values will be attempted. | false |
| volumeName | string | volumeName is the human-readable name of the StorageOS volume.  Volume<br/>names are only unique within a namespace. | false |
| volumeNamespace | string | volumeNamespace specifies the scope of the volume within StorageOS.  If no<br/>namespace is specified then the Pod's namespace will be used.  This allows the<br/>Kubernetes name scoping to be mirrored within StorageOS for tighter integration.<br/>Set VolumeName to any name to override the default behaviour.<br/>Set to "default" if you are not using namespaces within StorageOS.<br/>Namespaces that do not pre-exist within StorageOS will be created. | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].storageos.secretRef

secretRef specifies the secret to use for obtaining the StorageOS API<br/>credentials.  If not specified, default values will be attempted.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.openwebui.pipelines.volumes[index].vsphereVolume

vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine.<br/>Deprecated: VsphereVolume is deprecated. All operations for the in-tree vsphereVolume type<br/>are redirected to the csi.vsphere.vmware.com CSI driver.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| volumePath | string | volumePath is the path that identifies vSphere volume vmdk | true |
| fsType | string | fsType is filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. | false |
| storagePolicyID | string | storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName. | false |
| storagePolicyName | string | storagePolicyName is the storage Policy Based Management (SPBM) profile name. | false |
### LMDeployment.spec.openwebui.redis

Redis defines the Redis configuration for OpenWebUI

| Name | Type | Description | Required |
|------|------|-------------|----------|
| enabled | boolean | Enabled determines if Redis should be deployed automatically<br/>If false and RedisURL is not provided, Redis will not be deployed | false |
| image | string | Image is the Redis container image to use (including tag) | false |
| password | string | Password is the Redis password (optional) | false |
| [persistence](#lmdeploymentspecopenwebuiredispersistence) | object | Persistence defines Redis persistence configuration | false |
| redisUrl | string | RedisURL is the Redis connection URL<br/>If not provided and Enabled is true, Redis will be deployed automatically<br/>Format: redis://host:port/db or rediss://host:port/db for TLS | false |
| [resources](#lmdeploymentspecopenwebuiredisresources) | object | Resources defines the resource requirements for Redis pods | false |
| [service](#lmdeploymentspecopenwebuiredisservice) | object | Service defines the service configuration for Redis | false |
### LMDeployment.spec.openwebui.redis.persistence

Persistence defines Redis persistence configuration

| Name | Type | Description | Required |
|------|------|-------------|----------|
| enabled | boolean | Enabled determines if Redis data should be persisted | false |
| size | string | Size is the size of the persistent volume | false |
| storageClass | string | StorageClass is the storage class to use for persistent volumes | false |
### LMDeployment.spec.openwebui.redis.resources

Resources defines the resource requirements for Redis pods

| Name | Type | Description | Required |
|------|------|-------------|----------|
| limits | map[string]int or string | Limits describes the maximum amount of compute resources allowed | false |
| requests | map[string]int or string | Requests describes the minimum amount of compute resources required | false |
### LMDeployment.spec.openwebui.redis.service

Service defines the service configuration for Redis

| Name | Type | Description | Required |
|------|------|-------------|----------|
| port | integer | Port is the port to expose the service<br/>*Format*: int32<br/>*Minimum*: 1<br/>*Maximum*: 65535<br/> | false |
| type | enum | Type is the type of service to expose<br/>*Enum*: ClusterIP, NodePort, LoadBalancer<br/> | false |
### LMDeployment.spec.openwebui.resources

Resources defines the resource requirements for OpenWebUI pods

| Name | Type | Description | Required |
|------|------|-------------|----------|
| limits | map[string]int or string | Limits describes the maximum amount of compute resources allowed | false |
| requests | map[string]int or string | Requests describes the minimum amount of compute resources required | false |
### LMDeployment.spec.openwebui.service

Service defines the service configuration for OpenWebUI

| Name | Type | Description | Required |
|------|------|-------------|----------|
| port | integer | Port is the port to expose the service<br/>*Format*: int32<br/>*Minimum*: 1<br/>*Maximum*: 65535<br/> | false |
| type | enum | Type is the type of service to expose<br/>*Enum*: ClusterIP, NodePort, LoadBalancer<br/> | false |
### LMDeployment.spec.tabby

Tabby defines the Tabby deployment configuration

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [affinity](#lmdeploymentspectabbyaffinity) | object | Affinity defines pod affinity and anti-affinity rules for Tabby pods | false |
| chatModel | string | ChatModel is the name of the model to use for chat functionality<br/>Must be one of the models specified in spec.ollama.models or spec.vllm.model | false |
| completionModel | string | CompletionModel is the name of the model to use for code completion<br/>Must be one of the models specified in spec.ollama.models or spec.vllm.model | false |
| configMapName | string | ConfigMapName is the name of the ConfigMap containing Tabby configuration | false |
| device | enum | Device specifies the device type for Tabby,<br/>*Enum*: cpu, cuda, rocm, metal, vulkan<br/>*Default*: cpu<br/> | false |
| enabled | boolean | Enabled determines if Tabby should be deployed | false |
| [envVars](#lmdeploymentspectabbyenvvarsindex) | []object | EnvVars defines environment variables for Tabby | false |
| image | string | Image is the Tabby container image to use (including tag) | false |
| [ingress](#lmdeploymentspectabbyingress) | object | Ingress defines the ingress configuration for Tabby | false |
| [persistence](#lmdeploymentspectabbypersistence) | object | Persistence defines Tabby persistence configuration | false |
| replicas | integer | Replicas is the number of Tabby pods to run<br/>*Format*: int32<br/>*Minimum*: 1<br/>*Maximum*: 5<br/> | false |
| [resources](#lmdeploymentspectabbyresources) | object | Resources defines the resource requirements for Tabby pods | false |
| [service](#lmdeploymentspectabbyservice) | object | Service defines the service configuration for Tabby | false |
| [volumeMounts](#lmdeploymentspectabbyvolumemountsindex) | []object | VolumeMounts defines volume mounts for Tabby | false |
| [volumes](#lmdeploymentspectabbyvolumesindex) | []object | Volumes defines volumes for Tabby | false |
### LMDeployment.spec.tabby.affinity

Affinity defines pod affinity and anti-affinity rules for Tabby pods

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [nodeAffinity](#lmdeploymentspectabbyaffinitynodeaffinity) | object | Describes node affinity scheduling rules for the pod. | false |
| [podAffinity](#lmdeploymentspectabbyaffinitypodaffinity) | object | Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)). | false |
| [podAntiAffinity](#lmdeploymentspectabbyaffinitypodantiaffinity) | object | Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)). | false |
### LMDeployment.spec.tabby.affinity.nodeAffinity

Describes node affinity scheduling rules for the pod.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [preferredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspectabbyaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindex) | []object | The scheduler will prefer to schedule pods to nodes that satisfy<br/>the affinity expressions specified by this field, but it may choose<br/>a node that violates one or more of the expressions. The node that is<br/>most preferred is the one with the greatest sum of weights, i.e.<br/>for each node that meets all of the scheduling requirements (resource<br/>request, requiredDuringScheduling affinity expressions, etc.),<br/>compute a sum by iterating through the elements of this field and adding<br/>"weight" to the sum if the node matches the corresponding matchExpressions; the<br/>node(s) with the highest sum are the most preferred. | false |
| [requiredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspectabbyaffinitynodeaffinityrequiredduringschedulingignoredduringexecution) | object | If the affinity requirements specified by this field are not met at<br/>scheduling time, the pod will not be scheduled onto the node.<br/>If the affinity requirements specified by this field cease to be met<br/>at some point during pod execution (e.g. due to an update), the system<br/>may or may not try to eventually evict the pod from its node. | false |
### LMDeployment.spec.tabby.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index]

An empty preferred scheduling term matches all objects with implicit weight 0<br/>(i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [preference](#lmdeploymentspectabbyaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindexpreference) | object | A node selector term, associated with the corresponding weight. | true |
| weight | integer | Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.<br/>*Format*: int32<br/> | true |
### LMDeployment.spec.tabby.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].preference

A node selector term, associated with the corresponding weight.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspectabbyaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindexpreferencematchexpressionsindex) | []object | A list of node selector requirements by node's labels. | false |
| [matchFields](#lmdeploymentspectabbyaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindexpreferencematchfieldsindex) | []object | A list of node selector requirements by node's fields. | false |
### LMDeployment.spec.tabby.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].preference.matchExpressions[index]

A node selector requirement is a selector that contains values, a key, and an operator<br/>that relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The label key that the selector applies to. | true |
| operator | string | Represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt. | true |
| values | []string | An array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. If the operator is Gt or Lt, the values<br/>array must have a single element, which will be interpreted as an integer.<br/>This array is replaced during a strategic merge patch. | false |
### LMDeployment.spec.tabby.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].preference.matchFields[index]

A node selector requirement is a selector that contains values, a key, and an operator<br/>that relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The label key that the selector applies to. | true |
| operator | string | Represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt. | true |
| values | []string | An array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. If the operator is Gt or Lt, the values<br/>array must have a single element, which will be interpreted as an integer.<br/>This array is replaced during a strategic merge patch. | false |
### LMDeployment.spec.tabby.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution

If the affinity requirements specified by this field are not met at<br/>scheduling time, the pod will not be scheduled onto the node.<br/>If the affinity requirements specified by this field cease to be met<br/>at some point during pod execution (e.g. due to an update), the system<br/>may or may not try to eventually evict the pod from its node.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [nodeSelectorTerms](#lmdeploymentspectabbyaffinitynodeaffinityrequiredduringschedulingignoredduringexecutionnodeselectortermsindex) | []object | Required. A list of node selector terms. The terms are ORed. | true |
### LMDeployment.spec.tabby.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[index]

A null or empty node selector term matches no objects. The requirements of<br/>them are ANDed.<br/>The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspectabbyaffinitynodeaffinityrequiredduringschedulingignoredduringexecutionnodeselectortermsindexmatchexpressionsindex) | []object | A list of node selector requirements by node's labels. | false |
| [matchFields](#lmdeploymentspectabbyaffinitynodeaffinityrequiredduringschedulingignoredduringexecutionnodeselectortermsindexmatchfieldsindex) | []object | A list of node selector requirements by node's fields. | false |
### LMDeployment.spec.tabby.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[index].matchExpressions[index]

A node selector requirement is a selector that contains values, a key, and an operator<br/>that relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The label key that the selector applies to. | true |
| operator | string | Represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt. | true |
| values | []string | An array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. If the operator is Gt or Lt, the values<br/>array must have a single element, which will be interpreted as an integer.<br/>This array is replaced during a strategic merge patch. | false |
### LMDeployment.spec.tabby.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[index].matchFields[index]

A node selector requirement is a selector that contains values, a key, and an operator<br/>that relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The label key that the selector applies to. | true |
| operator | string | Represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt. | true |
| values | []string | An array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. If the operator is Gt or Lt, the values<br/>array must have a single element, which will be interpreted as an integer.<br/>This array is replaced during a strategic merge patch. | false |
### LMDeployment.spec.tabby.affinity.podAffinity

Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [preferredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspectabbyaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindex) | []object | The scheduler will prefer to schedule pods to nodes that satisfy<br/>the affinity expressions specified by this field, but it may choose<br/>a node that violates one or more of the expressions. The node that is<br/>most preferred is the one with the greatest sum of weights, i.e.<br/>for each node that meets all of the scheduling requirements (resource<br/>request, requiredDuringScheduling affinity expressions, etc.),<br/>compute a sum by iterating through the elements of this field and adding<br/>"weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the<br/>node(s) with the highest sum are the most preferred. | false |
| [requiredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspectabbyaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindex) | []object | If the affinity requirements specified by this field are not met at<br/>scheduling time, the pod will not be scheduled onto the node.<br/>If the affinity requirements specified by this field cease to be met<br/>at some point during pod execution (e.g. due to a pod label update), the<br/>system may or may not try to eventually evict the pod from its node.<br/>When there are multiple elements, the lists of nodes corresponding to each<br/>podAffinityTerm are intersected, i.e. all terms must be satisfied. | false |
### LMDeployment.spec.tabby.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index]

The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [podAffinityTerm](#lmdeploymentspectabbyaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinityterm) | object | Required. A pod affinity term, associated with the corresponding weight. | true |
| weight | integer | weight associated with matching the corresponding podAffinityTerm,<br/>in the range 1-100.<br/>*Format*: int32<br/> | true |
### LMDeployment.spec.tabby.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm

Required. A pod affinity term, associated with the corresponding weight.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| topologyKey | string | This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching<br/>the labelSelector in the specified namespaces, where co-located is defined as running on a node<br/>whose value of the label with key topologyKey matches that of any node on which any of the<br/>selected pods is running.<br/>Empty topologyKey is not allowed. | true |
| [labelSelector](#lmdeploymentspectabbyaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselector) | object | A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods. | false |
| matchLabelKeys | []string | MatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both matchLabelKeys and labelSelector.<br/>Also, matchLabelKeys cannot be set when labelSelector isn't set. | false |
| mismatchLabelKeys | []string | MismatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.<br/>Also, mismatchLabelKeys cannot be set when labelSelector isn't set. | false |
| [namespaceSelector](#lmdeploymentspectabbyaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselector) | object | A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces. | false |
| namespaces | []string | namespaces specifies a static list of namespace names that the term applies to.<br/>The term is applied to the union of the namespaces listed in this field<br/>and the ones selected by namespaceSelector.<br/>null or empty namespaces list and null namespaceSelector means "this pod's namespace". | false |
### LMDeployment.spec.tabby.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector

A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspectabbyaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.tabby.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.tabby.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector

A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspectabbyaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.tabby.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.tabby.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index]

Defines a set of pods (namely those matching the labelSelector<br/>relative to the given namespace(s)) that this pod should be<br/>co-located (affinity) or not co-located (anti-affinity) with,<br/>where co-located is defined as running on a node whose value of<br/>the label with key &lt;topologyKey&gt; matches that of any node on which<br/>a pod of the set of pods is running

| Name | Type | Description | Required |
|------|------|-------------|----------|
| topologyKey | string | This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching<br/>the labelSelector in the specified namespaces, where co-located is defined as running on a node<br/>whose value of the label with key topologyKey matches that of any node on which any of the<br/>selected pods is running.<br/>Empty topologyKey is not allowed. | true |
| [labelSelector](#lmdeploymentspectabbyaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexlabelselector) | object | A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods. | false |
| matchLabelKeys | []string | MatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both matchLabelKeys and labelSelector.<br/>Also, matchLabelKeys cannot be set when labelSelector isn't set. | false |
| mismatchLabelKeys | []string | MismatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.<br/>Also, mismatchLabelKeys cannot be set when labelSelector isn't set. | false |
| [namespaceSelector](#lmdeploymentspectabbyaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselector) | object | A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces. | false |
| namespaces | []string | namespaces specifies a static list of namespace names that the term applies to.<br/>The term is applied to the union of the namespaces listed in this field<br/>and the ones selected by namespaceSelector.<br/>null or empty namespaces list and null namespaceSelector means "this pod's namespace". | false |
### LMDeployment.spec.tabby.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector

A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspectabbyaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexlabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.tabby.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.tabby.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector

A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspectabbyaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.tabby.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.tabby.affinity.podAntiAffinity

Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [preferredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspectabbyaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindex) | []object | The scheduler will prefer to schedule pods to nodes that satisfy<br/>the anti-affinity expressions specified by this field, but it may choose<br/>a node that violates one or more of the expressions. The node that is<br/>most preferred is the one with the greatest sum of weights, i.e.<br/>for each node that meets all of the scheduling requirements (resource<br/>request, requiredDuringScheduling anti-affinity expressions, etc.),<br/>compute a sum by iterating through the elements of this field and adding<br/>"weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the<br/>node(s) with the highest sum are the most preferred. | false |
| [requiredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspectabbyaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindex) | []object | If the anti-affinity requirements specified by this field are not met at<br/>scheduling time, the pod will not be scheduled onto the node.<br/>If the anti-affinity requirements specified by this field cease to be met<br/>at some point during pod execution (e.g. due to a pod label update), the<br/>system may or may not try to eventually evict the pod from its node.<br/>When there are multiple elements, the lists of nodes corresponding to each<br/>podAffinityTerm are intersected, i.e. all terms must be satisfied. | false |
### LMDeployment.spec.tabby.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index]

The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [podAffinityTerm](#lmdeploymentspectabbyaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinityterm) | object | Required. A pod affinity term, associated with the corresponding weight. | true |
| weight | integer | weight associated with matching the corresponding podAffinityTerm,<br/>in the range 1-100.<br/>*Format*: int32<br/> | true |
### LMDeployment.spec.tabby.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm

Required. A pod affinity term, associated with the corresponding weight.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| topologyKey | string | This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching<br/>the labelSelector in the specified namespaces, where co-located is defined as running on a node<br/>whose value of the label with key topologyKey matches that of any node on which any of the<br/>selected pods is running.<br/>Empty topologyKey is not allowed. | true |
| [labelSelector](#lmdeploymentspectabbyaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselector) | object | A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods. | false |
| matchLabelKeys | []string | MatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both matchLabelKeys and labelSelector.<br/>Also, matchLabelKeys cannot be set when labelSelector isn't set. | false |
| mismatchLabelKeys | []string | MismatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.<br/>Also, mismatchLabelKeys cannot be set when labelSelector isn't set. | false |
| [namespaceSelector](#lmdeploymentspectabbyaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselector) | object | A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces. | false |
| namespaces | []string | namespaces specifies a static list of namespace names that the term applies to.<br/>The term is applied to the union of the namespaces listed in this field<br/>and the ones selected by namespaceSelector.<br/>null or empty namespaces list and null namespaceSelector means "this pod's namespace". | false |
### LMDeployment.spec.tabby.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector

A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspectabbyaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.tabby.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.tabby.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector

A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspectabbyaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.tabby.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.tabby.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index]

Defines a set of pods (namely those matching the labelSelector<br/>relative to the given namespace(s)) that this pod should be<br/>co-located (affinity) or not co-located (anti-affinity) with,<br/>where co-located is defined as running on a node whose value of<br/>the label with key &lt;topologyKey&gt; matches that of any node on which<br/>a pod of the set of pods is running

| Name | Type | Description | Required |
|------|------|-------------|----------|
| topologyKey | string | This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching<br/>the labelSelector in the specified namespaces, where co-located is defined as running on a node<br/>whose value of the label with key topologyKey matches that of any node on which any of the<br/>selected pods is running.<br/>Empty topologyKey is not allowed. | true |
| [labelSelector](#lmdeploymentspectabbyaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexlabelselector) | object | A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods. | false |
| matchLabelKeys | []string | MatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both matchLabelKeys and labelSelector.<br/>Also, matchLabelKeys cannot be set when labelSelector isn't set. | false |
| mismatchLabelKeys | []string | MismatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.<br/>Also, mismatchLabelKeys cannot be set when labelSelector isn't set. | false |
| [namespaceSelector](#lmdeploymentspectabbyaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselector) | object | A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces. | false |
| namespaces | []string | namespaces specifies a static list of namespace names that the term applies to.<br/>The term is applied to the union of the namespaces listed in this field<br/>and the ones selected by namespaceSelector.<br/>null or empty namespaces list and null namespaceSelector means "this pod's namespace". | false |
### LMDeployment.spec.tabby.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector

A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspectabbyaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexlabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.tabby.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.tabby.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector

A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspectabbyaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.tabby.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.tabby.envVars[index]

EnvVar represents an environment variable present in a Container.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the environment variable. Must be a C_IDENTIFIER. | true |
| value | string | Variable references $(VAR_NAME) are expanded<br/>using the previously defined environment variables in the container and<br/>any service environment variables. If a variable cannot be resolved,<br/>the reference in the input string will be unchanged. Double $$ are reduced<br/>to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.<br/>"$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)".<br/>Escaped references will never be expanded, regardless of whether the variable<br/>exists or not.<br/>Defaults to "". | false |
| [valueFrom](#lmdeploymentspectabbyenvvarsindexvaluefrom) | object | Source for the environment variable's value. Cannot be used if value is not empty. | false |
### LMDeployment.spec.tabby.envVars[index].valueFrom

Source for the environment variable's value. Cannot be used if value is not empty.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [configMapKeyRef](#lmdeploymentspectabbyenvvarsindexvaluefromconfigmapkeyref) | object | Selects a key of a ConfigMap. | false |
| [fieldRef](#lmdeploymentspectabbyenvvarsindexvaluefromfieldref) | object | Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['&lt;KEY&gt;']`, `metadata.annotations['&lt;KEY&gt;']`,<br/>spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs. | false |
| [resourceFieldRef](#lmdeploymentspectabbyenvvarsindexvaluefromresourcefieldref) | object | Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported. | false |
| [secretKeyRef](#lmdeploymentspectabbyenvvarsindexvaluefromsecretkeyref) | object | Selects a key of a secret in the pod's namespace | false |
### LMDeployment.spec.tabby.envVars[index].valueFrom.configMapKeyRef

Selects a key of a ConfigMap.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The key to select. | true |
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
| optional | boolean | Specify whether the ConfigMap or its key must be defined | false |
### LMDeployment.spec.tabby.envVars[index].valueFrom.fieldRef

Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['&lt;KEY&gt;']`, `metadata.annotations['&lt;KEY&gt;']`,<br/>spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| fieldPath | string | Path of the field to select in the specified API version. | true |
| apiVersion | string | Version of the schema the FieldPath is written in terms of, defaults to "v1". | false |
### LMDeployment.spec.tabby.envVars[index].valueFrom.resourceFieldRef

Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| resource | string | Required: resource to select | true |
| containerName | string | Container name: required for volumes, optional for env vars | false |
| divisor | int or string | Specifies the output format of the exposed resources, defaults to "1" | false |
### LMDeployment.spec.tabby.envVars[index].valueFrom.secretKeyRef

Selects a key of a secret in the pod's namespace

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The key of the secret to select from.  Must be a valid secret key. | true |
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
| optional | boolean | Specify whether the Secret or its key must be defined | false |
### LMDeployment.spec.tabby.ingress

Ingress defines the ingress configuration for Tabby

| Name | Type | Description | Required |
|------|------|-------------|----------|
| annotations | map[string]string | Annotations are custom annotations for the Ingress | false |
| host | string | Host is the hostname for the Ingress | false |
### LMDeployment.spec.tabby.persistence

Persistence defines Tabby persistence configuration

| Name | Type | Description | Required |
|------|------|-------------|----------|
| enabled | boolean | Enabled determines if Tabby data should be persisted | false |
| size | string | Size is the size of the persistent volume | false |
| storageClass | string | StorageClass is the storage class to use for persistent volumes | false |
### LMDeployment.spec.tabby.resources

Resources defines the resource requirements for Tabby pods

| Name | Type | Description | Required |
|------|------|-------------|----------|
| limits | map[string]int or string | Limits describes the maximum amount of compute resources allowed | false |
| requests | map[string]int or string | Requests describes the minimum amount of compute resources required | false |
### LMDeployment.spec.tabby.service

Service defines the service configuration for Tabby

| Name | Type | Description | Required |
|------|------|-------------|----------|
| port | integer | Port is the port to expose the service<br/>*Format*: int32<br/>*Minimum*: 1<br/>*Maximum*: 65535<br/> | false |
| type | enum | Type is the type of service to expose<br/>*Enum*: ClusterIP, NodePort, LoadBalancer<br/> | false |
### LMDeployment.spec.tabby.volumeMounts[index]

VolumeMount describes a mounting of a Volume within a container.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| mountPath | string | Path within the container at which the volume should be mounted.  Must<br/>not contain ':'. | true |
| name | string | This must match the Name of a Volume. | true |
| mountPropagation | string | mountPropagation determines how mounts are propagated from the host<br/>to container and the other way around.<br/>When not set, MountPropagationNone is used.<br/>This field is beta in 1.10.<br/>When RecursiveReadOnly is set to IfPossible or to Enabled, MountPropagation must be None or unspecified<br/>(which defaults to None). | false |
| readOnly | boolean | Mounted read-only if true, read-write otherwise (false or unspecified).<br/>Defaults to false. | false |
| recursiveReadOnly | string | RecursiveReadOnly specifies whether read-only mounts should be handled<br/>recursively.<br/><br/>If ReadOnly is false, this field has no meaning and must be unspecified.<br/><br/>If ReadOnly is true, and this field is set to Disabled, the mount is not made<br/>recursively read-only.  If this field is set to IfPossible, the mount is made<br/>recursively read-only, if it is supported by the container runtime.  If this<br/>field is set to Enabled, the mount is made recursively read-only if it is<br/>supported by the container runtime, otherwise the pod will not be started and<br/>an error will be generated to indicate the reason.<br/><br/>If this field is set to IfPossible or Enabled, MountPropagation must be set to<br/>None (or be unspecified, which defaults to None).<br/><br/>If this field is not specified, it is treated as an equivalent of Disabled. | false |
| subPath | string | Path within the volume from which the container's volume should be mounted.<br/>Defaults to "" (volume's root). | false |
| subPathExpr | string | Expanded path within the volume from which the container's volume should be mounted.<br/>Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.<br/>Defaults to "" (volume's root).<br/>SubPathExpr and SubPath are mutually exclusive. | false |
### LMDeployment.spec.tabby.volumes[index]

Volume represents a named volume in a pod that may be accessed by any container in the pod.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | name of the volume.<br/>Must be a DNS_LABEL and unique within the pod.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names | true |
| [awsElasticBlockStore](#lmdeploymentspectabbyvolumesindexawselasticblockstore) | object | awsElasticBlockStore represents an AWS Disk resource that is attached to a<br/>kubelet's host machine and then exposed to the pod.<br/>Deprecated: AWSElasticBlockStore is deprecated. All operations for the in-tree<br/>awsElasticBlockStore type are redirected to the ebs.csi.aws.com CSI driver.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore | false |
| [azureDisk](#lmdeploymentspectabbyvolumesindexazuredisk) | object | azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.<br/>Deprecated: AzureDisk is deprecated. All operations for the in-tree azureDisk type<br/>are redirected to the disk.csi.azure.com CSI driver. | false |
| [azureFile](#lmdeploymentspectabbyvolumesindexazurefile) | object | azureFile represents an Azure File Service mount on the host and bind mount to the pod.<br/>Deprecated: AzureFile is deprecated. All operations for the in-tree azureFile type<br/>are redirected to the file.csi.azure.com CSI driver. | false |
| [cephfs](#lmdeploymentspectabbyvolumesindexcephfs) | object | cephFS represents a Ceph FS mount on the host that shares a pod's lifetime.<br/>Deprecated: CephFS is deprecated and the in-tree cephfs type is no longer supported. | false |
| [cinder](#lmdeploymentspectabbyvolumesindexcinder) | object | cinder represents a cinder volume attached and mounted on kubelets host machine.<br/>Deprecated: Cinder is deprecated. All operations for the in-tree cinder type<br/>are redirected to the cinder.csi.openstack.org CSI driver.<br/>More info: https://examples.k8s.io/mysql-cinder-pd/README.md | false |
| [configMap](#lmdeploymentspectabbyvolumesindexconfigmap) | object | configMap represents a configMap that should populate this volume | false |
| [csi](#lmdeploymentspectabbyvolumesindexcsi) | object | csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers. | false |
| [downwardAPI](#lmdeploymentspectabbyvolumesindexdownwardapi) | object | downwardAPI represents downward API about the pod that should populate this volume | false |
| [emptyDir](#lmdeploymentspectabbyvolumesindexemptydir) | object | emptyDir represents a temporary directory that shares a pod's lifetime.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir | false |
| [ephemeral](#lmdeploymentspectabbyvolumesindexephemeral) | object | ephemeral represents a volume that is handled by a cluster storage driver.<br/>The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts,<br/>and deleted when the pod is removed.<br/><br/>Use this if:<br/>a) the volume is only needed while the pod runs,<br/>b) features of normal volumes like restoring from snapshot or capacity<br/>   tracking are needed,<br/>c) the storage driver is specified through a storage class, and<br/>d) the storage driver supports dynamic volume provisioning through<br/>   a PersistentVolumeClaim (see EphemeralVolumeSource for more<br/>   information on the connection between this volume type<br/>   and PersistentVolumeClaim).<br/><br/>Use PersistentVolumeClaim or one of the vendor-specific<br/>APIs for volumes that persist for longer than the lifecycle<br/>of an individual pod.<br/><br/>Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to<br/>be used that way - see the documentation of the driver for<br/>more information.<br/><br/>A pod can use both types of ephemeral volumes and<br/>persistent volumes at the same time. | false |
| [fc](#lmdeploymentspectabbyvolumesindexfc) | object | fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod. | false |
| [flexVolume](#lmdeploymentspectabbyvolumesindexflexvolume) | object | flexVolume represents a generic volume resource that is<br/>provisioned/attached using an exec based plugin.<br/>Deprecated: FlexVolume is deprecated. Consider using a CSIDriver instead. | false |
| [flocker](#lmdeploymentspectabbyvolumesindexflocker) | object | flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running.<br/>Deprecated: Flocker is deprecated and the in-tree flocker type is no longer supported. | false |
| [gcePersistentDisk](#lmdeploymentspectabbyvolumesindexgcepersistentdisk) | object | gcePersistentDisk represents a GCE Disk resource that is attached to a<br/>kubelet's host machine and then exposed to the pod.<br/>Deprecated: GCEPersistentDisk is deprecated. All operations for the in-tree<br/>gcePersistentDisk type are redirected to the pd.csi.storage.gke.io CSI driver.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk | false |
| [gitRepo](#lmdeploymentspectabbyvolumesindexgitrepo) | object | gitRepo represents a git repository at a particular revision.<br/>Deprecated: GitRepo is deprecated. To provision a container with a git repo, mount an<br/>EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir<br/>into the Pod's container. | false |
| [glusterfs](#lmdeploymentspectabbyvolumesindexglusterfs) | object | glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.<br/>Deprecated: Glusterfs is deprecated and the in-tree glusterfs type is no longer supported.<br/>More info: https://examples.k8s.io/volumes/glusterfs/README.md | false |
| [hostPath](#lmdeploymentspectabbyvolumesindexhostpath) | object | hostPath represents a pre-existing file or directory on the host<br/>machine that is directly exposed to the container. This is generally<br/>used for system agents or other privileged things that are allowed<br/>to see the host machine. Most containers will NOT need this.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath | false |
| [image](#lmdeploymentspectabbyvolumesindeximage) | object | image represents an OCI object (a container image or artifact) pulled and mounted on the kubelet's host machine.<br/>The volume is resolved at pod startup depending on which PullPolicy value is provided:<br/><br/>- Always: the kubelet always attempts to pull the reference. Container creation will fail If the pull fails.<br/>- Never: the kubelet never pulls the reference and only uses a local image or artifact. Container creation will fail if the reference isn't present.<br/>- IfNotPresent: the kubelet pulls if the reference isn't already present on disk. Container creation will fail if the reference isn't present and the pull fails.<br/><br/>The volume gets re-resolved if the pod gets deleted and recreated, which means that new remote content will become available on pod recreation.<br/>A failure to resolve or pull the image during pod startup will block containers from starting and may add significant latency. Failures will be retried using normal volume backoff and will be reported on the pod reason and message.<br/>The types of objects that may be mounted by this volume are defined by the container runtime implementation on a host machine and at minimum must include all valid types supported by the container image field.<br/>The OCI object gets mounted in a single directory (spec.containers[*].volumeMounts.mountPath) by merging the manifest layers in the same way as for container images.<br/>The volume will be mounted read-only (ro) and non-executable files (noexec).<br/>Sub path mounts for containers are not supported (spec.containers[*].volumeMounts.subpath) before 1.33.<br/>The field spec.securityContext.fsGroupChangePolicy has no effect on this volume type. | false |
| [iscsi](#lmdeploymentspectabbyvolumesindexiscsi) | object | iscsi represents an ISCSI Disk resource that is attached to a<br/>kubelet's host machine and then exposed to the pod.<br/>More info: https://examples.k8s.io/volumes/iscsi/README.md | false |
| [nfs](#lmdeploymentspectabbyvolumesindexnfs) | object | nfs represents an NFS mount on the host that shares a pod's lifetime<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs | false |
| [persistentVolumeClaim](#lmdeploymentspectabbyvolumesindexpersistentvolumeclaim) | object | persistentVolumeClaimVolumeSource represents a reference to a<br/>PersistentVolumeClaim in the same namespace.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims | false |
| [photonPersistentDisk](#lmdeploymentspectabbyvolumesindexphotonpersistentdisk) | object | photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine.<br/>Deprecated: PhotonPersistentDisk is deprecated and the in-tree photonPersistentDisk type is no longer supported. | false |
| [portworxVolume](#lmdeploymentspectabbyvolumesindexportworxvolume) | object | portworxVolume represents a portworx volume attached and mounted on kubelets host machine.<br/>Deprecated: PortworxVolume is deprecated. All operations for the in-tree portworxVolume type<br/>are redirected to the pxd.portworx.com CSI driver when the CSIMigrationPortworx feature-gate<br/>is on. | false |
| [projected](#lmdeploymentspectabbyvolumesindexprojected) | object | projected items for all in one resources secrets, configmaps, and downward API | false |
| [quobyte](#lmdeploymentspectabbyvolumesindexquobyte) | object | quobyte represents a Quobyte mount on the host that shares a pod's lifetime.<br/>Deprecated: Quobyte is deprecated and the in-tree quobyte type is no longer supported. | false |
| [rbd](#lmdeploymentspectabbyvolumesindexrbd) | object | rbd represents a Rados Block Device mount on the host that shares a pod's lifetime.<br/>Deprecated: RBD is deprecated and the in-tree rbd type is no longer supported.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md | false |
| [scaleIO](#lmdeploymentspectabbyvolumesindexscaleio) | object | scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.<br/>Deprecated: ScaleIO is deprecated and the in-tree scaleIO type is no longer supported. | false |
| [secret](#lmdeploymentspectabbyvolumesindexsecret) | object | secret represents a secret that should populate this volume.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#secret | false |
| [storageos](#lmdeploymentspectabbyvolumesindexstorageos) | object | storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.<br/>Deprecated: StorageOS is deprecated and the in-tree storageos type is no longer supported. | false |
| [vsphereVolume](#lmdeploymentspectabbyvolumesindexvspherevolume) | object | vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine.<br/>Deprecated: VsphereVolume is deprecated. All operations for the in-tree vsphereVolume type<br/>are redirected to the csi.vsphere.vmware.com CSI driver. | false |
### LMDeployment.spec.tabby.volumes[index].awsElasticBlockStore

awsElasticBlockStore represents an AWS Disk resource that is attached to a<br/>kubelet's host machine and then exposed to the pod.<br/>Deprecated: AWSElasticBlockStore is deprecated. All operations for the in-tree<br/>awsElasticBlockStore type are redirected to the ebs.csi.aws.com CSI driver.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore

| Name | Type | Description | Required |
|------|------|-------------|----------|
| volumeID | string | volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume).<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore | true |
| fsType | string | fsType is the filesystem type of the volume that you want to mount.<br/>Tip: Ensure that the filesystem type is supported by the host operating system.<br/>Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore | false |
| partition | integer | partition is the partition in the volume that you want to mount.<br/>If omitted, the default is to mount by volume name.<br/>Examples: For volume /dev/sda1, you specify the partition as "1".<br/>Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).<br/>*Format*: int32<br/> | false |
| readOnly | boolean | readOnly value true will force the readOnly setting in VolumeMounts.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore | false |
### LMDeployment.spec.tabby.volumes[index].azureDisk

azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.<br/>Deprecated: AzureDisk is deprecated. All operations for the in-tree azureDisk type<br/>are redirected to the disk.csi.azure.com CSI driver.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| diskName | string | diskName is the Name of the data disk in the blob storage | true |
| diskURI | string | diskURI is the URI of data disk in the blob storage | true |
| cachingMode | string | cachingMode is the Host Caching mode: None, Read Only, Read Write. | false |
| fsType | string | fsType is Filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br/>*Default*: ext4<br/> | false |
| kind | string | kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared | false |
| readOnly | boolean | readOnly Defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts.<br/>*Default*: false<br/> | false |
### LMDeployment.spec.tabby.volumes[index].azureFile

azureFile represents an Azure File Service mount on the host and bind mount to the pod.<br/>Deprecated: AzureFile is deprecated. All operations for the in-tree azureFile type<br/>are redirected to the file.csi.azure.com CSI driver.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| secretName | string | secretName is the  name of secret that contains Azure Storage Account Name and Key | true |
| shareName | string | shareName is the azure share Name | true |
| readOnly | boolean | readOnly defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts. | false |
### LMDeployment.spec.tabby.volumes[index].cephfs

cephFS represents a Ceph FS mount on the host that shares a pod's lifetime.<br/>Deprecated: CephFS is deprecated and the in-tree cephfs type is no longer supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| monitors | []string | monitors is Required: Monitors is a collection of Ceph monitors<br/>More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | true |
| path | string | path is Optional: Used as the mounted root, rather than the full Ceph tree, default is / | false |
| readOnly | boolean | readOnly is Optional: Defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts.<br/>More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | false |
| secretFile | string | secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret<br/>More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | false |
| [secretRef](#lmdeploymentspectabbyvolumesindexcephfssecretref) | object | secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty.<br/>More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | false |
| user | string | user is optional: User is the rados user name, default is admin<br/>More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | false |
### LMDeployment.spec.tabby.volumes[index].cephfs.secretRef

secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty.<br/>More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.tabby.volumes[index].cinder

cinder represents a cinder volume attached and mounted on kubelets host machine.<br/>Deprecated: Cinder is deprecated. All operations for the in-tree cinder type<br/>are redirected to the cinder.csi.openstack.org CSI driver.<br/>More info: https://examples.k8s.io/mysql-cinder-pd/README.md

| Name | Type | Description | Required |
|------|------|-------------|----------|
| volumeID | string | volumeID used to identify the volume in cinder.<br/>More info: https://examples.k8s.io/mysql-cinder-pd/README.md | true |
| fsType | string | fsType is the filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br/>More info: https://examples.k8s.io/mysql-cinder-pd/README.md | false |
| readOnly | boolean | readOnly defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts.<br/>More info: https://examples.k8s.io/mysql-cinder-pd/README.md | false |
| [secretRef](#lmdeploymentspectabbyvolumesindexcindersecretref) | object | secretRef is optional: points to a secret object containing parameters used to connect<br/>to OpenStack. | false |
### LMDeployment.spec.tabby.volumes[index].cinder.secretRef

secretRef is optional: points to a secret object containing parameters used to connect<br/>to OpenStack.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.tabby.volumes[index].configMap

configMap represents a configMap that should populate this volume

| Name | Type | Description | Required |
|------|------|-------------|----------|
| defaultMode | integer | defaultMode is optional: mode bits used to set permissions on created files by default.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>Defaults to 0644.<br/>Directories within the path are not affected by this setting.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
| [items](#lmdeploymentspectabbyvolumesindexconfigmapitemsindex) | []object | items if unspecified, each key-value pair in the Data field of the referenced<br/>ConfigMap will be projected into the volume as a file whose name is the<br/>key and content is the value. If specified, the listed keys will be<br/>projected into the specified paths, and unlisted keys will not be<br/>present. If a key is specified which is not present in the ConfigMap,<br/>the volume setup will error unless it is marked optional. Paths must be<br/>relative and may not contain the '..' path or start with '..'. | false |
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
| optional | boolean | optional specify whether the ConfigMap or its keys must be defined | false |
### LMDeployment.spec.tabby.volumes[index].configMap.items[index]

Maps a string key to a path within a volume.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the key to project. | true |
| path | string | path is the relative path of the file to map the key to.<br/>May not be an absolute path.<br/>May not contain the path element '..'.<br/>May not start with the string '..'. | true |
| mode | integer | mode is Optional: mode bits used to set permissions on this file.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>If not specified, the volume defaultMode will be used.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
### LMDeployment.spec.tabby.volumes[index].csi

csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| driver | string | driver is the name of the CSI driver that handles this volume.<br/>Consult with your admin for the correct name as registered in the cluster. | true |
| fsType | string | fsType to mount. Ex. "ext4", "xfs", "ntfs".<br/>If not provided, the empty value is passed to the associated CSI driver<br/>which will determine the default filesystem to apply. | false |
| [nodePublishSecretRef](#lmdeploymentspectabbyvolumesindexcsinodepublishsecretref) | object | nodePublishSecretRef is a reference to the secret object containing<br/>sensitive information to pass to the CSI driver to complete the CSI<br/>NodePublishVolume and NodeUnpublishVolume calls.<br/>This field is optional, and  may be empty if no secret is required. If the<br/>secret object contains more than one secret, all secret references are passed. | false |
| readOnly | boolean | readOnly specifies a read-only configuration for the volume.<br/>Defaults to false (read/write). | false |
| volumeAttributes | map[string]string | volumeAttributes stores driver-specific properties that are passed to the CSI<br/>driver. Consult your driver's documentation for supported values. | false |
### LMDeployment.spec.tabby.volumes[index].csi.nodePublishSecretRef

nodePublishSecretRef is a reference to the secret object containing<br/>sensitive information to pass to the CSI driver to complete the CSI<br/>NodePublishVolume and NodeUnpublishVolume calls.<br/>This field is optional, and  may be empty if no secret is required. If the<br/>secret object contains more than one secret, all secret references are passed.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.tabby.volumes[index].downwardAPI

downwardAPI represents downward API about the pod that should populate this volume

| Name | Type | Description | Required |
|------|------|-------------|----------|
| defaultMode | integer | Optional: mode bits to use on created files by default. Must be a<br/>Optional: mode bits used to set permissions on created files by default.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>Defaults to 0644.<br/>Directories within the path are not affected by this setting.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
| [items](#lmdeploymentspectabbyvolumesindexdownwardapiitemsindex) | []object | Items is a list of downward API volume file | false |
### LMDeployment.spec.tabby.volumes[index].downwardAPI.items[index]

DownwardAPIVolumeFile represents information to create the file containing the pod field

| Name | Type | Description | Required |
|------|------|-------------|----------|
| path | string | Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..' | true |
| [fieldRef](#lmdeploymentspectabbyvolumesindexdownwardapiitemsindexfieldref) | object | Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported. | false |
| mode | integer | Optional: mode bits used to set permissions on this file, must be an octal value<br/>between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>If not specified, the volume defaultMode will be used.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
| [resourceFieldRef](#lmdeploymentspectabbyvolumesindexdownwardapiitemsindexresourcefieldref) | object | Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported. | false |
### LMDeployment.spec.tabby.volumes[index].downwardAPI.items[index].fieldRef

Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| fieldPath | string | Path of the field to select in the specified API version. | true |
| apiVersion | string | Version of the schema the FieldPath is written in terms of, defaults to "v1". | false |
### LMDeployment.spec.tabby.volumes[index].downwardAPI.items[index].resourceFieldRef

Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| resource | string | Required: resource to select | true |
| containerName | string | Container name: required for volumes, optional for env vars | false |
| divisor | int or string | Specifies the output format of the exposed resources, defaults to "1" | false |
### LMDeployment.spec.tabby.volumes[index].emptyDir

emptyDir represents a temporary directory that shares a pod's lifetime.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir

| Name | Type | Description | Required |
|------|------|-------------|----------|
| medium | string | medium represents what type of storage medium should back this directory.<br/>The default is "" which means to use the node's default medium.<br/>Must be an empty string (default) or Memory.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir | false |
| sizeLimit | int or string | sizeLimit is the total amount of local storage required for this EmptyDir volume.<br/>The size limit is also applicable for memory medium.<br/>The maximum usage on memory medium EmptyDir would be the minimum value between<br/>the SizeLimit specified here and the sum of memory limits of all containers in a pod.<br/>The default is nil which means that the limit is undefined.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir | false |
### LMDeployment.spec.tabby.volumes[index].ephemeral

ephemeral represents a volume that is handled by a cluster storage driver.<br/>The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts,<br/>and deleted when the pod is removed.<br/><br/>Use this if:<br/>a) the volume is only needed while the pod runs,<br/>b) features of normal volumes like restoring from snapshot or capacity<br/>   tracking are needed,<br/>c) the storage driver is specified through a storage class, and<br/>d) the storage driver supports dynamic volume provisioning through<br/>   a PersistentVolumeClaim (see EphemeralVolumeSource for more<br/>   information on the connection between this volume type<br/>   and PersistentVolumeClaim).<br/><br/>Use PersistentVolumeClaim or one of the vendor-specific<br/>APIs for volumes that persist for longer than the lifecycle<br/>of an individual pod.<br/><br/>Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to<br/>be used that way - see the documentation of the driver for<br/>more information.<br/><br/>A pod can use both types of ephemeral volumes and<br/>persistent volumes at the same time.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [volumeClaimTemplate](#lmdeploymentspectabbyvolumesindexephemeralvolumeclaimtemplate) | object | Will be used to create a stand-alone PVC to provision the volume.<br/>The pod in which this EphemeralVolumeSource is embedded will be the<br/>owner of the PVC, i.e. the PVC will be deleted together with the<br/>pod.  The name of the PVC will be `&lt;pod name&gt;-&lt;volume name&gt;` where<br/>`&lt;volume name&gt;` is the name from the `PodSpec.Volumes` array<br/>entry. Pod validation will reject the pod if the concatenated name<br/>is not valid for a PVC (for example, too long).<br/><br/>An existing PVC with that name that is not owned by the pod<br/>will *not* be used for the pod to avoid using an unrelated<br/>volume by mistake. Starting the pod is then blocked until<br/>the unrelated PVC is removed. If such a pre-created PVC is<br/>meant to be used by the pod, the PVC has to updated with an<br/>owner reference to the pod once the pod exists. Normally<br/>this should not be necessary, but it may be useful when<br/>manually reconstructing a broken cluster.<br/><br/>This field is read-only and no changes will be made by Kubernetes<br/>to the PVC after it has been created.<br/><br/>Required, must not be nil. | false |
### LMDeployment.spec.tabby.volumes[index].ephemeral.volumeClaimTemplate

Will be used to create a stand-alone PVC to provision the volume.<br/>The pod in which this EphemeralVolumeSource is embedded will be the<br/>owner of the PVC, i.e. the PVC will be deleted together with the<br/>pod.  The name of the PVC will be `&lt;pod name&gt;-&lt;volume name&gt;` where<br/>`&lt;volume name&gt;` is the name from the `PodSpec.Volumes` array<br/>entry. Pod validation will reject the pod if the concatenated name<br/>is not valid for a PVC (for example, too long).<br/><br/>An existing PVC with that name that is not owned by the pod<br/>will *not* be used for the pod to avoid using an unrelated<br/>volume by mistake. Starting the pod is then blocked until<br/>the unrelated PVC is removed. If such a pre-created PVC is<br/>meant to be used by the pod, the PVC has to updated with an<br/>owner reference to the pod once the pod exists. Normally<br/>this should not be necessary, but it may be useful when<br/>manually reconstructing a broken cluster.<br/><br/>This field is read-only and no changes will be made by Kubernetes<br/>to the PVC after it has been created.<br/><br/>Required, must not be nil.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [spec](#lmdeploymentspectabbyvolumesindexephemeralvolumeclaimtemplatespec) | object | The specification for the PersistentVolumeClaim. The entire content is<br/>copied unchanged into the PVC that gets created from this<br/>template. The same fields as in a PersistentVolumeClaim<br/>are also valid here. | true |
| metadata | object | May contain labels and annotations that will be copied into the PVC<br/>when creating it. No other fields are allowed and will be rejected during<br/>validation. | false |
### LMDeployment.spec.tabby.volumes[index].ephemeral.volumeClaimTemplate.spec

The specification for the PersistentVolumeClaim. The entire content is<br/>copied unchanged into the PVC that gets created from this<br/>template. The same fields as in a PersistentVolumeClaim<br/>are also valid here.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| accessModes | []string | accessModes contains the desired access modes the volume should have.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 | false |
| [dataSource](#lmdeploymentspectabbyvolumesindexephemeralvolumeclaimtemplatespecdatasource) | object | dataSource field can be used to specify either:<br/>* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)<br/>* An existing PVC (PersistentVolumeClaim)<br/>If the provisioner or an external controller can support the specified data source,<br/>it will create a new volume based on the contents of the specified data source.<br/>When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,<br/>and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.<br/>If the namespace is specified, then dataSourceRef will not be copied to dataSource. | false |
| [dataSourceRef](#lmdeploymentspectabbyvolumesindexephemeralvolumeclaimtemplatespecdatasourceref) | object | dataSourceRef specifies the object from which to populate the volume with data, if a non-empty<br/>volume is desired. This may be any object from a non-empty API group (non<br/>core object) or a PersistentVolumeClaim object.<br/>When this field is specified, volume binding will only succeed if the type of<br/>the specified object matches some installed volume populator or dynamic<br/>provisioner.<br/>This field will replace the functionality of the dataSource field and as such<br/>if both fields are non-empty, they must have the same value. For backwards<br/>compatibility, when namespace isn't specified in dataSourceRef,<br/>both fields (dataSource and dataSourceRef) will be set to the same<br/>value automatically if one of them is empty and the other is non-empty.<br/>When namespace is specified in dataSourceRef,<br/>dataSource isn't set to the same value and must be empty.<br/>There are three important differences between dataSource and dataSourceRef:<br/>* While dataSource only allows two specific types of objects, dataSourceRef<br/>  allows any non-core object, as well as PersistentVolumeClaim objects.<br/>* While dataSource ignores disallowed values (dropping them), dataSourceRef<br/>  preserves all values, and generates an error if a disallowed value is<br/>  specified.<br/>* While dataSource only allows local objects, dataSourceRef allows objects<br/>  in any namespaces.<br/>(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.<br/>(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled. | false |
| [resources](#lmdeploymentspectabbyvolumesindexephemeralvolumeclaimtemplatespecresources) | object | resources represents the minimum resources the volume should have.<br/>If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements<br/>that are lower than previous value but must still be higher than capacity recorded in the<br/>status field of the claim.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources | false |
| [selector](#lmdeploymentspectabbyvolumesindexephemeralvolumeclaimtemplatespecselector) | object | selector is a label query over volumes to consider for binding. | false |
| storageClassName | string | storageClassName is the name of the StorageClass required by the claim.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1 | false |
| volumeAttributesClassName | string | volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.<br/>If specified, the CSI driver will create or update the volume with the attributes defined<br/>in the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,<br/>it can be changed after the claim is created. An empty string value means that no VolumeAttributesClass<br/>will be applied to the claim but it's not allowed to reset this field to empty string once it is set.<br/>If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClass<br/>will be set by the persistentvolume controller if it exists.<br/>If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will be<br/>set to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resource<br/>exists.<br/>More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/<br/>(Beta) Using this field requires the VolumeAttributesClass feature gate to be enabled (off by default). | false |
| volumeMode | string | volumeMode defines what type of volume is required by the claim.<br/>Value of Filesystem is implied when not included in claim spec. | false |
| volumeName | string | volumeName is the binding reference to the PersistentVolume backing this claim. | false |
### LMDeployment.spec.tabby.volumes[index].ephemeral.volumeClaimTemplate.spec.dataSource

dataSource field can be used to specify either:<br/>* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)<br/>* An existing PVC (PersistentVolumeClaim)<br/>If the provisioner or an external controller can support the specified data source,<br/>it will create a new volume based on the contents of the specified data source.<br/>When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,<br/>and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.<br/>If the namespace is specified, then dataSourceRef will not be copied to dataSource.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| kind | string | Kind is the type of resource being referenced | true |
| name | string | Name is the name of resource being referenced | true |
| apiGroup | string | APIGroup is the group for the resource being referenced.<br/>If APIGroup is not specified, the specified Kind must be in the core API group.<br/>For any other third-party types, APIGroup is required. | false |
### LMDeployment.spec.tabby.volumes[index].ephemeral.volumeClaimTemplate.spec.dataSourceRef

dataSourceRef specifies the object from which to populate the volume with data, if a non-empty<br/>volume is desired. This may be any object from a non-empty API group (non<br/>core object) or a PersistentVolumeClaim object.<br/>When this field is specified, volume binding will only succeed if the type of<br/>the specified object matches some installed volume populator or dynamic<br/>provisioner.<br/>This field will replace the functionality of the dataSource field and as such<br/>if both fields are non-empty, they must have the same value. For backwards<br/>compatibility, when namespace isn't specified in dataSourceRef,<br/>both fields (dataSource and dataSourceRef) will be set to the same<br/>value automatically if one of them is empty and the other is non-empty.<br/>When namespace is specified in dataSourceRef,<br/>dataSource isn't set to the same value and must be empty.<br/>There are three important differences between dataSource and dataSourceRef:<br/>* While dataSource only allows two specific types of objects, dataSourceRef<br/>  allows any non-core object, as well as PersistentVolumeClaim objects.<br/>* While dataSource ignores disallowed values (dropping them), dataSourceRef<br/>  preserves all values, and generates an error if a disallowed value is<br/>  specified.<br/>* While dataSource only allows local objects, dataSourceRef allows objects<br/>  in any namespaces.<br/>(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.<br/>(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| kind | string | Kind is the type of resource being referenced | true |
| name | string | Name is the name of resource being referenced | true |
| apiGroup | string | APIGroup is the group for the resource being referenced.<br/>If APIGroup is not specified, the specified Kind must be in the core API group.<br/>For any other third-party types, APIGroup is required. | false |
| namespace | string | Namespace is the namespace of resource being referenced<br/>Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.<br/>(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled. | false |
### LMDeployment.spec.tabby.volumes[index].ephemeral.volumeClaimTemplate.spec.resources

resources represents the minimum resources the volume should have.<br/>If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements<br/>that are lower than previous value but must still be higher than capacity recorded in the<br/>status field of the claim.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources

| Name | Type | Description | Required |
|------|------|-------------|----------|
| limits | map[string]int or string | Limits describes the maximum amount of compute resources allowed.<br/>More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ | false |
| requests | map[string]int or string | Requests describes the minimum amount of compute resources required.<br/>If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,<br/>otherwise to an implementation-defined value. Requests cannot exceed Limits.<br/>More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ | false |
### LMDeployment.spec.tabby.volumes[index].ephemeral.volumeClaimTemplate.spec.selector

selector is a label query over volumes to consider for binding.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspectabbyvolumesindexephemeralvolumeclaimtemplatespecselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.tabby.volumes[index].ephemeral.volumeClaimTemplate.spec.selector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.tabby.volumes[index].fc

fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| fsType | string | fsType is the filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. | false |
| lun | integer | lun is Optional: FC target lun number<br/>*Format*: int32<br/> | false |
| readOnly | boolean | readOnly is Optional: Defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts. | false |
| targetWWNs | []string | targetWWNs is Optional: FC target worldwide names (WWNs) | false |
| wwids | []string | wwids Optional: FC volume world wide identifiers (wwids)<br/>Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously. | false |
### LMDeployment.spec.tabby.volumes[index].flexVolume

flexVolume represents a generic volume resource that is<br/>provisioned/attached using an exec based plugin.<br/>Deprecated: FlexVolume is deprecated. Consider using a CSIDriver instead.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| driver | string | driver is the name of the driver to use for this volume. | true |
| fsType | string | fsType is the filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script. | false |
| options | map[string]string | options is Optional: this field holds extra command options if any. | false |
| readOnly | boolean | readOnly is Optional: defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts. | false |
| [secretRef](#lmdeploymentspectabbyvolumesindexflexvolumesecretref) | object | secretRef is Optional: secretRef is reference to the secret object containing<br/>sensitive information to pass to the plugin scripts. This may be<br/>empty if no secret object is specified. If the secret object<br/>contains more than one secret, all secrets are passed to the plugin<br/>scripts. | false |
### LMDeployment.spec.tabby.volumes[index].flexVolume.secretRef

secretRef is Optional: secretRef is reference to the secret object containing<br/>sensitive information to pass to the plugin scripts. This may be<br/>empty if no secret object is specified. If the secret object<br/>contains more than one secret, all secrets are passed to the plugin<br/>scripts.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.tabby.volumes[index].flocker

flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running.<br/>Deprecated: Flocker is deprecated and the in-tree flocker type is no longer supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| datasetName | string | datasetName is Name of the dataset stored as metadata -&gt; name on the dataset for Flocker<br/>should be considered as deprecated | false |
| datasetUUID | string | datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset | false |
### LMDeployment.spec.tabby.volumes[index].gcePersistentDisk

gcePersistentDisk represents a GCE Disk resource that is attached to a<br/>kubelet's host machine and then exposed to the pod.<br/>Deprecated: GCEPersistentDisk is deprecated. All operations for the in-tree<br/>gcePersistentDisk type are redirected to the pd.csi.storage.gke.io CSI driver.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk

| Name | Type | Description | Required |
|------|------|-------------|----------|
| pdName | string | pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk | true |
| fsType | string | fsType is filesystem type of the volume that you want to mount.<br/>Tip: Ensure that the filesystem type is supported by the host operating system.<br/>Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk | false |
| partition | integer | partition is the partition in the volume that you want to mount.<br/>If omitted, the default is to mount by volume name.<br/>Examples: For volume /dev/sda1, you specify the partition as "1".<br/>Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk<br/>*Format*: int32<br/> | false |
| readOnly | boolean | readOnly here will force the ReadOnly setting in VolumeMounts.<br/>Defaults to false.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk | false |
### LMDeployment.spec.tabby.volumes[index].gitRepo

gitRepo represents a git repository at a particular revision.<br/>Deprecated: GitRepo is deprecated. To provision a container with a git repo, mount an<br/>EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir<br/>into the Pod's container.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| repository | string | repository is the URL | true |
| directory | string | directory is the target directory name.<br/>Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the<br/>git repository.  Otherwise, if specified, the volume will contain the git repository in<br/>the subdirectory with the given name. | false |
| revision | string | revision is the commit hash for the specified revision. | false |
### LMDeployment.spec.tabby.volumes[index].glusterfs

glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.<br/>Deprecated: Glusterfs is deprecated and the in-tree glusterfs type is no longer supported.<br/>More info: https://examples.k8s.io/volumes/glusterfs/README.md

| Name | Type | Description | Required |
|------|------|-------------|----------|
| endpoints | string | endpoints is the endpoint name that details Glusterfs topology.<br/>More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod | true |
| path | string | path is the Glusterfs volume path.<br/>More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod | true |
| readOnly | boolean | readOnly here will force the Glusterfs volume to be mounted with read-only permissions.<br/>Defaults to false.<br/>More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod | false |
### LMDeployment.spec.tabby.volumes[index].hostPath

hostPath represents a pre-existing file or directory on the host<br/>machine that is directly exposed to the container. This is generally<br/>used for system agents or other privileged things that are allowed<br/>to see the host machine. Most containers will NOT need this.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath

| Name | Type | Description | Required |
|------|------|-------------|----------|
| path | string | path of the directory on the host.<br/>If the path is a symlink, it will follow the link to the real path.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath | true |
| type | string | type for HostPath Volume<br/>Defaults to ""<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath | false |
### LMDeployment.spec.tabby.volumes[index].image

image represents an OCI object (a container image or artifact) pulled and mounted on the kubelet's host machine.<br/>The volume is resolved at pod startup depending on which PullPolicy value is provided:<br/><br/>- Always: the kubelet always attempts to pull the reference. Container creation will fail If the pull fails.<br/>- Never: the kubelet never pulls the reference and only uses a local image or artifact. Container creation will fail if the reference isn't present.<br/>- IfNotPresent: the kubelet pulls if the reference isn't already present on disk. Container creation will fail if the reference isn't present and the pull fails.<br/><br/>The volume gets re-resolved if the pod gets deleted and recreated, which means that new remote content will become available on pod recreation.<br/>A failure to resolve or pull the image during pod startup will block containers from starting and may add significant latency. Failures will be retried using normal volume backoff and will be reported on the pod reason and message.<br/>The types of objects that may be mounted by this volume are defined by the container runtime implementation on a host machine and at minimum must include all valid types supported by the container image field.<br/>The OCI object gets mounted in a single directory (spec.containers[*].volumeMounts.mountPath) by merging the manifest layers in the same way as for container images.<br/>The volume will be mounted read-only (ro) and non-executable files (noexec).<br/>Sub path mounts for containers are not supported (spec.containers[*].volumeMounts.subpath) before 1.33.<br/>The field spec.securityContext.fsGroupChangePolicy has no effect on this volume type.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| pullPolicy | string | Policy for pulling OCI objects. Possible values are:<br/>Always: the kubelet always attempts to pull the reference. Container creation will fail If the pull fails.<br/>Never: the kubelet never pulls the reference and only uses a local image or artifact. Container creation will fail if the reference isn't present.<br/>IfNotPresent: the kubelet pulls if the reference isn't already present on disk. Container creation will fail if the reference isn't present and the pull fails.<br/>Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. | false |
| reference | string | Required: Image or artifact reference to be used.<br/>Behaves in the same way as pod.spec.containers[*].image.<br/>Pull secrets will be assembled in the same way as for the container image by looking up node credentials, SA image pull secrets, and pod spec image pull secrets.<br/>More info: https://kubernetes.io/docs/concepts/containers/images<br/>This field is optional to allow higher level config management to default or override<br/>container images in workload controllers like Deployments and StatefulSets. | false |
### LMDeployment.spec.tabby.volumes[index].iscsi

iscsi represents an ISCSI Disk resource that is attached to a<br/>kubelet's host machine and then exposed to the pod.<br/>More info: https://examples.k8s.io/volumes/iscsi/README.md

| Name | Type | Description | Required |
|------|------|-------------|----------|
| iqn | string | iqn is the target iSCSI Qualified Name. | true |
| lun | integer | lun represents iSCSI Target Lun number.<br/>*Format*: int32<br/> | true |
| targetPortal | string | targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port<br/>is other than default (typically TCP ports 860 and 3260). | true |
| chapAuthDiscovery | boolean | chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication | false |
| chapAuthSession | boolean | chapAuthSession defines whether support iSCSI Session CHAP authentication | false |
| fsType | string | fsType is the filesystem type of the volume that you want to mount.<br/>Tip: Ensure that the filesystem type is supported by the host operating system.<br/>Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi | false |
| initiatorName | string | initiatorName is the custom iSCSI Initiator Name.<br/>If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface<br/>&lt;target portal&gt;:&lt;volume name&gt; will be created for the connection. | false |
| iscsiInterface | string | iscsiInterface is the interface Name that uses an iSCSI transport.<br/>Defaults to 'default' (tcp).<br/>*Default*: default<br/> | false |
| portals | []string | portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port<br/>is other than default (typically TCP ports 860 and 3260). | false |
| readOnly | boolean | readOnly here will force the ReadOnly setting in VolumeMounts.<br/>Defaults to false. | false |
| [secretRef](#lmdeploymentspectabbyvolumesindexiscsisecretref) | object | secretRef is the CHAP Secret for iSCSI target and initiator authentication | false |
### LMDeployment.spec.tabby.volumes[index].iscsi.secretRef

secretRef is the CHAP Secret for iSCSI target and initiator authentication

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.tabby.volumes[index].nfs

nfs represents an NFS mount on the host that shares a pod's lifetime<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs

| Name | Type | Description | Required |
|------|------|-------------|----------|
| path | string | path that is exported by the NFS server.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs | true |
| server | string | server is the hostname or IP address of the NFS server.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs | true |
| readOnly | boolean | readOnly here will force the NFS export to be mounted with read-only permissions.<br/>Defaults to false.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs | false |
### LMDeployment.spec.tabby.volumes[index].persistentVolumeClaim

persistentVolumeClaimVolumeSource represents a reference to a<br/>PersistentVolumeClaim in the same namespace.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims

| Name | Type | Description | Required |
|------|------|-------------|----------|
| claimName | string | claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims | true |
| readOnly | boolean | readOnly Will force the ReadOnly setting in VolumeMounts.<br/>Default false. | false |
### LMDeployment.spec.tabby.volumes[index].photonPersistentDisk

photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine.<br/>Deprecated: PhotonPersistentDisk is deprecated and the in-tree photonPersistentDisk type is no longer supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| pdID | string | pdID is the ID that identifies Photon Controller persistent disk | true |
| fsType | string | fsType is the filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. | false |
### LMDeployment.spec.tabby.volumes[index].portworxVolume

portworxVolume represents a portworx volume attached and mounted on kubelets host machine.<br/>Deprecated: PortworxVolume is deprecated. All operations for the in-tree portworxVolume type<br/>are redirected to the pxd.portworx.com CSI driver when the CSIMigrationPortworx feature-gate<br/>is on.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| volumeID | string | volumeID uniquely identifies a Portworx volume | true |
| fsType | string | fSType represents the filesystem type to mount<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified. | false |
| readOnly | boolean | readOnly defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts. | false |
### LMDeployment.spec.tabby.volumes[index].projected

projected items for all in one resources secrets, configmaps, and downward API

| Name | Type | Description | Required |
|------|------|-------------|----------|
| defaultMode | integer | defaultMode are the mode bits used to set permissions on created files by default.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>Directories within the path are not affected by this setting.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
| [sources](#lmdeploymentspectabbyvolumesindexprojectedsourcesindex) | []object | sources is the list of volume projections. Each entry in this list<br/>handles one source. | false |
### LMDeployment.spec.tabby.volumes[index].projected.sources[index]

Projection that may be projected along with other supported volume types.<br/>Exactly one of these fields must be set.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [clusterTrustBundle](#lmdeploymentspectabbyvolumesindexprojectedsourcesindexclustertrustbundle) | object | ClusterTrustBundle allows a pod to access the `.spec.trustBundle` field<br/>of ClusterTrustBundle objects in an auto-updating file.<br/><br/>Alpha, gated by the ClusterTrustBundleProjection feature gate.<br/><br/>ClusterTrustBundle objects can either be selected by name, or by the<br/>combination of signer name and a label selector.<br/><br/>Kubelet performs aggressive normalization of the PEM contents written<br/>into the pod filesystem.  Esoteric PEM features such as inter-block<br/>comments and block headers are stripped.  Certificates are deduplicated.<br/>The ordering of certificates within the file is arbitrary, and Kubelet<br/>may change the order over time. | false |
| [configMap](#lmdeploymentspectabbyvolumesindexprojectedsourcesindexconfigmap) | object | configMap information about the configMap data to project | false |
| [downwardAPI](#lmdeploymentspectabbyvolumesindexprojectedsourcesindexdownwardapi) | object | downwardAPI information about the downwardAPI data to project | false |
| [secret](#lmdeploymentspectabbyvolumesindexprojectedsourcesindexsecret) | object | secret information about the secret data to project | false |
| [serviceAccountToken](#lmdeploymentspectabbyvolumesindexprojectedsourcesindexserviceaccounttoken) | object | serviceAccountToken is information about the serviceAccountToken data to project | false |
### LMDeployment.spec.tabby.volumes[index].projected.sources[index].clusterTrustBundle

ClusterTrustBundle allows a pod to access the `.spec.trustBundle` field<br/>of ClusterTrustBundle objects in an auto-updating file.<br/><br/>Alpha, gated by the ClusterTrustBundleProjection feature gate.<br/><br/>ClusterTrustBundle objects can either be selected by name, or by the<br/>combination of signer name and a label selector.<br/><br/>Kubelet performs aggressive normalization of the PEM contents written<br/>into the pod filesystem.  Esoteric PEM features such as inter-block<br/>comments and block headers are stripped.  Certificates are deduplicated.<br/>The ordering of certificates within the file is arbitrary, and Kubelet<br/>may change the order over time.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| path | string | Relative path from the volume root to write the bundle. | true |
| [labelSelector](#lmdeploymentspectabbyvolumesindexprojectedsourcesindexclustertrustbundlelabelselector) | object | Select all ClusterTrustBundles that match this label selector.  Only has<br/>effect if signerName is set.  Mutually-exclusive with name.  If unset,<br/>interpreted as "match nothing".  If set but empty, interpreted as "match<br/>everything". | false |
| name | string | Select a single ClusterTrustBundle by object name.  Mutually-exclusive<br/>with signerName and labelSelector. | false |
| optional | boolean | If true, don't block pod startup if the referenced ClusterTrustBundle(s)<br/>aren't available.  If using name, then the named ClusterTrustBundle is<br/>allowed not to exist.  If using signerName, then the combination of<br/>signerName and labelSelector is allowed to match zero<br/>ClusterTrustBundles. | false |
| signerName | string | Select all ClusterTrustBundles that match this signer name.<br/>Mutually-exclusive with name.  The contents of all selected<br/>ClusterTrustBundles will be unified and deduplicated. | false |
### LMDeployment.spec.tabby.volumes[index].projected.sources[index].clusterTrustBundle.labelSelector

Select all ClusterTrustBundles that match this label selector.  Only has<br/>effect if signerName is set.  Mutually-exclusive with name.  If unset,<br/>interpreted as "match nothing".  If set but empty, interpreted as "match<br/>everything".

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspectabbyvolumesindexprojectedsourcesindexclustertrustbundlelabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.tabby.volumes[index].projected.sources[index].clusterTrustBundle.labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.tabby.volumes[index].projected.sources[index].configMap

configMap information about the configMap data to project

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [items](#lmdeploymentspectabbyvolumesindexprojectedsourcesindexconfigmapitemsindex) | []object | items if unspecified, each key-value pair in the Data field of the referenced<br/>ConfigMap will be projected into the volume as a file whose name is the<br/>key and content is the value. If specified, the listed keys will be<br/>projected into the specified paths, and unlisted keys will not be<br/>present. If a key is specified which is not present in the ConfigMap,<br/>the volume setup will error unless it is marked optional. Paths must be<br/>relative and may not contain the '..' path or start with '..'. | false |
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
| optional | boolean | optional specify whether the ConfigMap or its keys must be defined | false |
### LMDeployment.spec.tabby.volumes[index].projected.sources[index].configMap.items[index]

Maps a string key to a path within a volume.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the key to project. | true |
| path | string | path is the relative path of the file to map the key to.<br/>May not be an absolute path.<br/>May not contain the path element '..'.<br/>May not start with the string '..'. | true |
| mode | integer | mode is Optional: mode bits used to set permissions on this file.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>If not specified, the volume defaultMode will be used.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
### LMDeployment.spec.tabby.volumes[index].projected.sources[index].downwardAPI

downwardAPI information about the downwardAPI data to project

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [items](#lmdeploymentspectabbyvolumesindexprojectedsourcesindexdownwardapiitemsindex) | []object | Items is a list of DownwardAPIVolume file | false |
### LMDeployment.spec.tabby.volumes[index].projected.sources[index].downwardAPI.items[index]

DownwardAPIVolumeFile represents information to create the file containing the pod field

| Name | Type | Description | Required |
|------|------|-------------|----------|
| path | string | Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..' | true |
| [fieldRef](#lmdeploymentspectabbyvolumesindexprojectedsourcesindexdownwardapiitemsindexfieldref) | object | Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported. | false |
| mode | integer | Optional: mode bits used to set permissions on this file, must be an octal value<br/>between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>If not specified, the volume defaultMode will be used.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
| [resourceFieldRef](#lmdeploymentspectabbyvolumesindexprojectedsourcesindexdownwardapiitemsindexresourcefieldref) | object | Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported. | false |
### LMDeployment.spec.tabby.volumes[index].projected.sources[index].downwardAPI.items[index].fieldRef

Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| fieldPath | string | Path of the field to select in the specified API version. | true |
| apiVersion | string | Version of the schema the FieldPath is written in terms of, defaults to "v1". | false |
### LMDeployment.spec.tabby.volumes[index].projected.sources[index].downwardAPI.items[index].resourceFieldRef

Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| resource | string | Required: resource to select | true |
| containerName | string | Container name: required for volumes, optional for env vars | false |
| divisor | int or string | Specifies the output format of the exposed resources, defaults to "1" | false |
### LMDeployment.spec.tabby.volumes[index].projected.sources[index].secret

secret information about the secret data to project

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [items](#lmdeploymentspectabbyvolumesindexprojectedsourcesindexsecretitemsindex) | []object | items if unspecified, each key-value pair in the Data field of the referenced<br/>Secret will be projected into the volume as a file whose name is the<br/>key and content is the value. If specified, the listed keys will be<br/>projected into the specified paths, and unlisted keys will not be<br/>present. If a key is specified which is not present in the Secret,<br/>the volume setup will error unless it is marked optional. Paths must be<br/>relative and may not contain the '..' path or start with '..'. | false |
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
| optional | boolean | optional field specify whether the Secret or its key must be defined | false |
### LMDeployment.spec.tabby.volumes[index].projected.sources[index].secret.items[index]

Maps a string key to a path within a volume.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the key to project. | true |
| path | string | path is the relative path of the file to map the key to.<br/>May not be an absolute path.<br/>May not contain the path element '..'.<br/>May not start with the string '..'. | true |
| mode | integer | mode is Optional: mode bits used to set permissions on this file.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>If not specified, the volume defaultMode will be used.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
### LMDeployment.spec.tabby.volumes[index].projected.sources[index].serviceAccountToken

serviceAccountToken is information about the serviceAccountToken data to project

| Name | Type | Description | Required |
|------|------|-------------|----------|
| path | string | path is the path relative to the mount point of the file to project the<br/>token into. | true |
| audience | string | audience is the intended audience of the token. A recipient of a token<br/>must identify itself with an identifier specified in the audience of the<br/>token, and otherwise should reject the token. The audience defaults to the<br/>identifier of the apiserver. | false |
| expirationSeconds | integer | expirationSeconds is the requested duration of validity of the service<br/>account token. As the token approaches expiration, the kubelet volume<br/>plugin will proactively rotate the service account token. The kubelet will<br/>start trying to rotate the token if the token is older than 80 percent of<br/>its time to live or if the token is older than 24 hours.Defaults to 1 hour<br/>and must be at least 10 minutes.<br/>*Format*: int64<br/> | false |
### LMDeployment.spec.tabby.volumes[index].quobyte

quobyte represents a Quobyte mount on the host that shares a pod's lifetime.<br/>Deprecated: Quobyte is deprecated and the in-tree quobyte type is no longer supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| registry | string | registry represents a single or multiple Quobyte Registry services<br/>specified as a string as host:port pair (multiple entries are separated with commas)<br/>which acts as the central registry for volumes | true |
| volume | string | volume is a string that references an already created Quobyte volume by name. | true |
| group | string | group to map volume access to<br/>Default is no group | false |
| readOnly | boolean | readOnly here will force the Quobyte volume to be mounted with read-only permissions.<br/>Defaults to false. | false |
| tenant | string | tenant owning the given Quobyte volume in the Backend<br/>Used with dynamically provisioned Quobyte volumes, value is set by the plugin | false |
| user | string | user to map volume access to<br/>Defaults to serivceaccount user | false |
### LMDeployment.spec.tabby.volumes[index].rbd

rbd represents a Rados Block Device mount on the host that shares a pod's lifetime.<br/>Deprecated: RBD is deprecated and the in-tree rbd type is no longer supported.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md

| Name | Type | Description | Required |
|------|------|-------------|----------|
| image | string | image is the rados image name.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | true |
| monitors | []string | monitors is a collection of Ceph monitors.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | true |
| fsType | string | fsType is the filesystem type of the volume that you want to mount.<br/>Tip: Ensure that the filesystem type is supported by the host operating system.<br/>Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd | false |
| keyring | string | keyring is the path to key ring for RBDUser.<br/>Default is /etc/ceph/keyring.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it<br/>*Default*: /etc/ceph/keyring<br/> | false |
| pool | string | pool is the rados pool name.<br/>Default is rbd.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it<br/>*Default*: rbd<br/> | false |
| readOnly | boolean | readOnly here will force the ReadOnly setting in VolumeMounts.<br/>Defaults to false.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | false |
| [secretRef](#lmdeploymentspectabbyvolumesindexrbdsecretref) | object | secretRef is name of the authentication secret for RBDUser. If provided<br/>overrides keyring.<br/>Default is nil.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | false |
| user | string | user is the rados user name.<br/>Default is admin.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it<br/>*Default*: admin<br/> | false |
### LMDeployment.spec.tabby.volumes[index].rbd.secretRef

secretRef is name of the authentication secret for RBDUser. If provided<br/>overrides keyring.<br/>Default is nil.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.tabby.volumes[index].scaleIO

scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.<br/>Deprecated: ScaleIO is deprecated and the in-tree scaleIO type is no longer supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| gateway | string | gateway is the host address of the ScaleIO API Gateway. | true |
| [secretRef](#lmdeploymentspectabbyvolumesindexscaleiosecretref) | object | secretRef references to the secret for ScaleIO user and other<br/>sensitive information. If this is not provided, Login operation will fail. | true |
| system | string | system is the name of the storage system as configured in ScaleIO. | true |
| fsType | string | fsType is the filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs".<br/>Default is "xfs".<br/>*Default*: xfs<br/> | false |
| protectionDomain | string | protectionDomain is the name of the ScaleIO Protection Domain for the configured storage. | false |
| readOnly | boolean | readOnly Defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts. | false |
| sslEnabled | boolean | sslEnabled Flag enable/disable SSL communication with Gateway, default false | false |
| storageMode | string | storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.<br/>Default is ThinProvisioned.<br/>*Default*: ThinProvisioned<br/> | false |
| storagePool | string | storagePool is the ScaleIO Storage Pool associated with the protection domain. | false |
| volumeName | string | volumeName is the name of a volume already created in the ScaleIO system<br/>that is associated with this volume source. | false |
### LMDeployment.spec.tabby.volumes[index].scaleIO.secretRef

secretRef references to the secret for ScaleIO user and other<br/>sensitive information. If this is not provided, Login operation will fail.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.tabby.volumes[index].secret

secret represents a secret that should populate this volume.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#secret

| Name | Type | Description | Required |
|------|------|-------------|----------|
| defaultMode | integer | defaultMode is Optional: mode bits used to set permissions on created files by default.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values<br/>for mode bits. Defaults to 0644.<br/>Directories within the path are not affected by this setting.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
| [items](#lmdeploymentspectabbyvolumesindexsecretitemsindex) | []object | items If unspecified, each key-value pair in the Data field of the referenced<br/>Secret will be projected into the volume as a file whose name is the<br/>key and content is the value. If specified, the listed keys will be<br/>projected into the specified paths, and unlisted keys will not be<br/>present. If a key is specified which is not present in the Secret,<br/>the volume setup will error unless it is marked optional. Paths must be<br/>relative and may not contain the '..' path or start with '..'. | false |
| optional | boolean | optional field specify whether the Secret or its keys must be defined | false |
| secretName | string | secretName is the name of the secret in the pod's namespace to use.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#secret | false |
### LMDeployment.spec.tabby.volumes[index].secret.items[index]

Maps a string key to a path within a volume.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the key to project. | true |
| path | string | path is the relative path of the file to map the key to.<br/>May not be an absolute path.<br/>May not contain the path element '..'.<br/>May not start with the string '..'. | true |
| mode | integer | mode is Optional: mode bits used to set permissions on this file.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>If not specified, the volume defaultMode will be used.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
### LMDeployment.spec.tabby.volumes[index].storageos

storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.<br/>Deprecated: StorageOS is deprecated and the in-tree storageos type is no longer supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| fsType | string | fsType is the filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. | false |
| readOnly | boolean | readOnly defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts. | false |
| [secretRef](#lmdeploymentspectabbyvolumesindexstorageossecretref) | object | secretRef specifies the secret to use for obtaining the StorageOS API<br/>credentials.  If not specified, default values will be attempted. | false |
| volumeName | string | volumeName is the human-readable name of the StorageOS volume.  Volume<br/>names are only unique within a namespace. | false |
| volumeNamespace | string | volumeNamespace specifies the scope of the volume within StorageOS.  If no<br/>namespace is specified then the Pod's namespace will be used.  This allows the<br/>Kubernetes name scoping to be mirrored within StorageOS for tighter integration.<br/>Set VolumeName to any name to override the default behaviour.<br/>Set to "default" if you are not using namespaces within StorageOS.<br/>Namespaces that do not pre-exist within StorageOS will be created. | false |
### LMDeployment.spec.tabby.volumes[index].storageos.secretRef

secretRef specifies the secret to use for obtaining the StorageOS API<br/>credentials.  If not specified, default values will be attempted.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.tabby.volumes[index].vsphereVolume

vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine.<br/>Deprecated: VsphereVolume is deprecated. All operations for the in-tree vsphereVolume type<br/>are redirected to the csi.vsphere.vmware.com CSI driver.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| volumePath | string | volumePath is the path that identifies vSphere volume vmdk | true |
| fsType | string | fsType is filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. | false |
| storagePolicyID | string | storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName. | false |
| storagePolicyName | string | storagePolicyName is the storage Policy Based Management (SPBM) profile name. | false |
### LMDeployment.spec.vllm

VLLM defines the vLLM deployment configuration

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [affinity](#lmdeploymentspecvllmaffinity) | object | Affinity defines pod affinity and anti-affinity rules for vLLM pods | false |
| enabled | boolean | Enabled determines if vLLM should be deployed instead of Ollama | false |
| [envVars](#lmdeploymentspecvllmenvvarsindex) | []object | EnvVars defines environment variables for vLLM | false |
| image | string | Image is the vLLM container image to use (including tag) | false |
| model | string |  | false |
| [persistence](#lmdeploymentspecvllmpersistence) | object | Persistence defines vLLM persistence configuration | false |
| replicas | integer | Replicas is the number of vLLM pods to run<br/>*Format*: int32<br/>*Minimum*: 1<br/>*Maximum*: 10<br/> | false |
| [resources](#lmdeploymentspecvllmresources) | object | Resources defines the resource requirements for vLLM pods | false |
| [service](#lmdeploymentspecvllmservice) | object | Service defines the service configuration for vLLM | false |
| [volumeMounts](#lmdeploymentspecvllmvolumemountsindex) | []object | VolumeMounts defines volume mounts for vLLM | false |
| [volumes](#lmdeploymentspecvllmvolumesindex) | []object | Volumes defines volumes for vLLM | false |
### LMDeployment.spec.vllm.affinity

Affinity defines pod affinity and anti-affinity rules for vLLM pods

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [nodeAffinity](#lmdeploymentspecvllmaffinitynodeaffinity) | object | Describes node affinity scheduling rules for the pod. | false |
| [podAffinity](#lmdeploymentspecvllmaffinitypodaffinity) | object | Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)). | false |
| [podAntiAffinity](#lmdeploymentspecvllmaffinitypodantiaffinity) | object | Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)). | false |
### LMDeployment.spec.vllm.affinity.nodeAffinity

Describes node affinity scheduling rules for the pod.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [preferredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspecvllmaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindex) | []object | The scheduler will prefer to schedule pods to nodes that satisfy<br/>the affinity expressions specified by this field, but it may choose<br/>a node that violates one or more of the expressions. The node that is<br/>most preferred is the one with the greatest sum of weights, i.e.<br/>for each node that meets all of the scheduling requirements (resource<br/>request, requiredDuringScheduling affinity expressions, etc.),<br/>compute a sum by iterating through the elements of this field and adding<br/>"weight" to the sum if the node matches the corresponding matchExpressions; the<br/>node(s) with the highest sum are the most preferred. | false |
| [requiredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspecvllmaffinitynodeaffinityrequiredduringschedulingignoredduringexecution) | object | If the affinity requirements specified by this field are not met at<br/>scheduling time, the pod will not be scheduled onto the node.<br/>If the affinity requirements specified by this field cease to be met<br/>at some point during pod execution (e.g. due to an update), the system<br/>may or may not try to eventually evict the pod from its node. | false |
### LMDeployment.spec.vllm.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index]

An empty preferred scheduling term matches all objects with implicit weight 0<br/>(i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [preference](#lmdeploymentspecvllmaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindexpreference) | object | A node selector term, associated with the corresponding weight. | true |
| weight | integer | Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.<br/>*Format*: int32<br/> | true |
### LMDeployment.spec.vllm.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].preference

A node selector term, associated with the corresponding weight.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecvllmaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindexpreferencematchexpressionsindex) | []object | A list of node selector requirements by node's labels. | false |
| [matchFields](#lmdeploymentspecvllmaffinitynodeaffinitypreferredduringschedulingignoredduringexecutionindexpreferencematchfieldsindex) | []object | A list of node selector requirements by node's fields. | false |
### LMDeployment.spec.vllm.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].preference.matchExpressions[index]

A node selector requirement is a selector that contains values, a key, and an operator<br/>that relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The label key that the selector applies to. | true |
| operator | string | Represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt. | true |
| values | []string | An array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. If the operator is Gt or Lt, the values<br/>array must have a single element, which will be interpreted as an integer.<br/>This array is replaced during a strategic merge patch. | false |
### LMDeployment.spec.vllm.affinity.nodeAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].preference.matchFields[index]

A node selector requirement is a selector that contains values, a key, and an operator<br/>that relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The label key that the selector applies to. | true |
| operator | string | Represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt. | true |
| values | []string | An array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. If the operator is Gt or Lt, the values<br/>array must have a single element, which will be interpreted as an integer.<br/>This array is replaced during a strategic merge patch. | false |
### LMDeployment.spec.vllm.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution

If the affinity requirements specified by this field are not met at<br/>scheduling time, the pod will not be scheduled onto the node.<br/>If the affinity requirements specified by this field cease to be met<br/>at some point during pod execution (e.g. due to an update), the system<br/>may or may not try to eventually evict the pod from its node.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [nodeSelectorTerms](#lmdeploymentspecvllmaffinitynodeaffinityrequiredduringschedulingignoredduringexecutionnodeselectortermsindex) | []object | Required. A list of node selector terms. The terms are ORed. | true |
### LMDeployment.spec.vllm.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[index]

A null or empty node selector term matches no objects. The requirements of<br/>them are ANDed.<br/>The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecvllmaffinitynodeaffinityrequiredduringschedulingignoredduringexecutionnodeselectortermsindexmatchexpressionsindex) | []object | A list of node selector requirements by node's labels. | false |
| [matchFields](#lmdeploymentspecvllmaffinitynodeaffinityrequiredduringschedulingignoredduringexecutionnodeselectortermsindexmatchfieldsindex) | []object | A list of node selector requirements by node's fields. | false |
### LMDeployment.spec.vllm.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[index].matchExpressions[index]

A node selector requirement is a selector that contains values, a key, and an operator<br/>that relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The label key that the selector applies to. | true |
| operator | string | Represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt. | true |
| values | []string | An array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. If the operator is Gt or Lt, the values<br/>array must have a single element, which will be interpreted as an integer.<br/>This array is replaced during a strategic merge patch. | false |
### LMDeployment.spec.vllm.affinity.nodeAffinity.requiredDuringSchedulingIgnoredDuringExecution.nodeSelectorTerms[index].matchFields[index]

A node selector requirement is a selector that contains values, a key, and an operator<br/>that relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The label key that the selector applies to. | true |
| operator | string | Represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt. | true |
| values | []string | An array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. If the operator is Gt or Lt, the values<br/>array must have a single element, which will be interpreted as an integer.<br/>This array is replaced during a strategic merge patch. | false |
### LMDeployment.spec.vllm.affinity.podAffinity

Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [preferredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspecvllmaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindex) | []object | The scheduler will prefer to schedule pods to nodes that satisfy<br/>the affinity expressions specified by this field, but it may choose<br/>a node that violates one or more of the expressions. The node that is<br/>most preferred is the one with the greatest sum of weights, i.e.<br/>for each node that meets all of the scheduling requirements (resource<br/>request, requiredDuringScheduling affinity expressions, etc.),<br/>compute a sum by iterating through the elements of this field and adding<br/>"weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the<br/>node(s) with the highest sum are the most preferred. | false |
| [requiredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspecvllmaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindex) | []object | If the affinity requirements specified by this field are not met at<br/>scheduling time, the pod will not be scheduled onto the node.<br/>If the affinity requirements specified by this field cease to be met<br/>at some point during pod execution (e.g. due to a pod label update), the<br/>system may or may not try to eventually evict the pod from its node.<br/>When there are multiple elements, the lists of nodes corresponding to each<br/>podAffinityTerm are intersected, i.e. all terms must be satisfied. | false |
### LMDeployment.spec.vllm.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index]

The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [podAffinityTerm](#lmdeploymentspecvllmaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinityterm) | object | Required. A pod affinity term, associated with the corresponding weight. | true |
| weight | integer | weight associated with matching the corresponding podAffinityTerm,<br/>in the range 1-100.<br/>*Format*: int32<br/> | true |
### LMDeployment.spec.vllm.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm

Required. A pod affinity term, associated with the corresponding weight.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| topologyKey | string | This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching<br/>the labelSelector in the specified namespaces, where co-located is defined as running on a node<br/>whose value of the label with key topologyKey matches that of any node on which any of the<br/>selected pods is running.<br/>Empty topologyKey is not allowed. | true |
| [labelSelector](#lmdeploymentspecvllmaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselector) | object | A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods. | false |
| matchLabelKeys | []string | MatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both matchLabelKeys and labelSelector.<br/>Also, matchLabelKeys cannot be set when labelSelector isn't set. | false |
| mismatchLabelKeys | []string | MismatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.<br/>Also, mismatchLabelKeys cannot be set when labelSelector isn't set. | false |
| [namespaceSelector](#lmdeploymentspecvllmaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselector) | object | A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces. | false |
| namespaces | []string | namespaces specifies a static list of namespace names that the term applies to.<br/>The term is applied to the union of the namespaces listed in this field<br/>and the ones selected by namespaceSelector.<br/>null or empty namespaces list and null namespaceSelector means "this pod's namespace". | false |
### LMDeployment.spec.vllm.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector

A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecvllmaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.vllm.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.vllm.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector

A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecvllmaffinitypodaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.vllm.affinity.podAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.vllm.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index]

Defines a set of pods (namely those matching the labelSelector<br/>relative to the given namespace(s)) that this pod should be<br/>co-located (affinity) or not co-located (anti-affinity) with,<br/>where co-located is defined as running on a node whose value of<br/>the label with key &lt;topologyKey&gt; matches that of any node on which<br/>a pod of the set of pods is running

| Name | Type | Description | Required |
|------|------|-------------|----------|
| topologyKey | string | This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching<br/>the labelSelector in the specified namespaces, where co-located is defined as running on a node<br/>whose value of the label with key topologyKey matches that of any node on which any of the<br/>selected pods is running.<br/>Empty topologyKey is not allowed. | true |
| [labelSelector](#lmdeploymentspecvllmaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexlabelselector) | object | A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods. | false |
| matchLabelKeys | []string | MatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both matchLabelKeys and labelSelector.<br/>Also, matchLabelKeys cannot be set when labelSelector isn't set. | false |
| mismatchLabelKeys | []string | MismatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.<br/>Also, mismatchLabelKeys cannot be set when labelSelector isn't set. | false |
| [namespaceSelector](#lmdeploymentspecvllmaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselector) | object | A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces. | false |
| namespaces | []string | namespaces specifies a static list of namespace names that the term applies to.<br/>The term is applied to the union of the namespaces listed in this field<br/>and the ones selected by namespaceSelector.<br/>null or empty namespaces list and null namespaceSelector means "this pod's namespace". | false |
### LMDeployment.spec.vllm.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector

A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecvllmaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexlabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.vllm.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.vllm.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector

A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecvllmaffinitypodaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.vllm.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.vllm.affinity.podAntiAffinity

Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [preferredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspecvllmaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindex) | []object | The scheduler will prefer to schedule pods to nodes that satisfy<br/>the anti-affinity expressions specified by this field, but it may choose<br/>a node that violates one or more of the expressions. The node that is<br/>most preferred is the one with the greatest sum of weights, i.e.<br/>for each node that meets all of the scheduling requirements (resource<br/>request, requiredDuringScheduling anti-affinity expressions, etc.),<br/>compute a sum by iterating through the elements of this field and adding<br/>"weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the<br/>node(s) with the highest sum are the most preferred. | false |
| [requiredDuringSchedulingIgnoredDuringExecution](#lmdeploymentspecvllmaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindex) | []object | If the anti-affinity requirements specified by this field are not met at<br/>scheduling time, the pod will not be scheduled onto the node.<br/>If the anti-affinity requirements specified by this field cease to be met<br/>at some point during pod execution (e.g. due to a pod label update), the<br/>system may or may not try to eventually evict the pod from its node.<br/>When there are multiple elements, the lists of nodes corresponding to each<br/>podAffinityTerm are intersected, i.e. all terms must be satisfied. | false |
### LMDeployment.spec.vllm.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index]

The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [podAffinityTerm](#lmdeploymentspecvllmaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinityterm) | object | Required. A pod affinity term, associated with the corresponding weight. | true |
| weight | integer | weight associated with matching the corresponding podAffinityTerm,<br/>in the range 1-100.<br/>*Format*: int32<br/> | true |
### LMDeployment.spec.vllm.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm

Required. A pod affinity term, associated with the corresponding weight.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| topologyKey | string | This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching<br/>the labelSelector in the specified namespaces, where co-located is defined as running on a node<br/>whose value of the label with key topologyKey matches that of any node on which any of the<br/>selected pods is running.<br/>Empty topologyKey is not allowed. | true |
| [labelSelector](#lmdeploymentspecvllmaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselector) | object | A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods. | false |
| matchLabelKeys | []string | MatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both matchLabelKeys and labelSelector.<br/>Also, matchLabelKeys cannot be set when labelSelector isn't set. | false |
| mismatchLabelKeys | []string | MismatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.<br/>Also, mismatchLabelKeys cannot be set when labelSelector isn't set. | false |
| [namespaceSelector](#lmdeploymentspecvllmaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselector) | object | A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces. | false |
| namespaces | []string | namespaces specifies a static list of namespace names that the term applies to.<br/>The term is applied to the union of the namespaces listed in this field<br/>and the ones selected by namespaceSelector.<br/>null or empty namespaces list and null namespaceSelector means "this pod's namespace". | false |
### LMDeployment.spec.vllm.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector

A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecvllmaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermlabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.vllm.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.vllm.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector

A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecvllmaffinitypodantiaffinitypreferredduringschedulingignoredduringexecutionindexpodaffinitytermnamespaceselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.vllm.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution[index].podAffinityTerm.namespaceSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.vllm.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index]

Defines a set of pods (namely those matching the labelSelector<br/>relative to the given namespace(s)) that this pod should be<br/>co-located (affinity) or not co-located (anti-affinity) with,<br/>where co-located is defined as running on a node whose value of<br/>the label with key &lt;topologyKey&gt; matches that of any node on which<br/>a pod of the set of pods is running

| Name | Type | Description | Required |
|------|------|-------------|----------|
| topologyKey | string | This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching<br/>the labelSelector in the specified namespaces, where co-located is defined as running on a node<br/>whose value of the label with key topologyKey matches that of any node on which any of the<br/>selected pods is running.<br/>Empty topologyKey is not allowed. | true |
| [labelSelector](#lmdeploymentspecvllmaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexlabelselector) | object | A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods. | false |
| matchLabelKeys | []string | MatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key in (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both matchLabelKeys and labelSelector.<br/>Also, matchLabelKeys cannot be set when labelSelector isn't set. | false |
| mismatchLabelKeys | []string | MismatchLabelKeys is a set of pod label keys to select which pods will<br/>be taken into consideration. The keys are used to lookup values from the<br/>incoming pod labels, those key-value labels are merged with `labelSelector` as `key notin (value)`<br/>to select the group of existing pods which pods will be taken into consideration<br/>for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming<br/>pod labels will be ignored. The default value is empty.<br/>The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.<br/>Also, mismatchLabelKeys cannot be set when labelSelector isn't set. | false |
| [namespaceSelector](#lmdeploymentspecvllmaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselector) | object | A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces. | false |
| namespaces | []string | namespaces specifies a static list of namespace names that the term applies to.<br/>The term is applied to the union of the namespaces listed in this field<br/>and the ones selected by namespaceSelector.<br/>null or empty namespaces list and null namespaceSelector means "this pod's namespace". | false |
### LMDeployment.spec.vllm.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector

A label query over a set of resources, in this case pods.<br/>If it's null, this PodAffinityTerm matches with no Pods.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecvllmaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexlabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.vllm.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.vllm.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector

A label query over the set of namespaces that the term applies to.<br/>The term is applied to the union of the namespaces selected by this field<br/>and the ones listed in the namespaces field.<br/>null selector and null or empty namespaces list means "this pod's namespace".<br/>An empty selector (&#123;&#125;) matches all namespaces.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecvllmaffinitypodantiaffinityrequiredduringschedulingignoredduringexecutionindexnamespaceselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.vllm.affinity.podAntiAffinity.requiredDuringSchedulingIgnoredDuringExecution[index].namespaceSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.vllm.envVars[index]

EnvVar represents an environment variable present in a Container.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the environment variable. Must be a C_IDENTIFIER. | true |
| value | string | Variable references $(VAR_NAME) are expanded<br/>using the previously defined environment variables in the container and<br/>any service environment variables. If a variable cannot be resolved,<br/>the reference in the input string will be unchanged. Double $$ are reduced<br/>to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.<br/>"$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)".<br/>Escaped references will never be expanded, regardless of whether the variable<br/>exists or not.<br/>Defaults to "". | false |
| [valueFrom](#lmdeploymentspecvllmenvvarsindexvaluefrom) | object | Source for the environment variable's value. Cannot be used if value is not empty. | false |
### LMDeployment.spec.vllm.envVars[index].valueFrom

Source for the environment variable's value. Cannot be used if value is not empty.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [configMapKeyRef](#lmdeploymentspecvllmenvvarsindexvaluefromconfigmapkeyref) | object | Selects a key of a ConfigMap. | false |
| [fieldRef](#lmdeploymentspecvllmenvvarsindexvaluefromfieldref) | object | Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['&lt;KEY&gt;']`, `metadata.annotations['&lt;KEY&gt;']`,<br/>spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs. | false |
| [resourceFieldRef](#lmdeploymentspecvllmenvvarsindexvaluefromresourcefieldref) | object | Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported. | false |
| [secretKeyRef](#lmdeploymentspecvllmenvvarsindexvaluefromsecretkeyref) | object | Selects a key of a secret in the pod's namespace | false |
### LMDeployment.spec.vllm.envVars[index].valueFrom.configMapKeyRef

Selects a key of a ConfigMap.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The key to select. | true |
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
| optional | boolean | Specify whether the ConfigMap or its key must be defined | false |
### LMDeployment.spec.vllm.envVars[index].valueFrom.fieldRef

Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['&lt;KEY&gt;']`, `metadata.annotations['&lt;KEY&gt;']`,<br/>spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| fieldPath | string | Path of the field to select in the specified API version. | true |
| apiVersion | string | Version of the schema the FieldPath is written in terms of, defaults to "v1". | false |
### LMDeployment.spec.vllm.envVars[index].valueFrom.resourceFieldRef

Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| resource | string | Required: resource to select | true |
| containerName | string | Container name: required for volumes, optional for env vars | false |
| divisor | int or string | Specifies the output format of the exposed resources, defaults to "1" | false |
### LMDeployment.spec.vllm.envVars[index].valueFrom.secretKeyRef

Selects a key of a secret in the pod's namespace

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | The key of the secret to select from.  Must be a valid secret key. | true |
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
| optional | boolean | Specify whether the Secret or its key must be defined | false |
### LMDeployment.spec.vllm.persistence

Persistence defines vLLM persistence configuration

| Name | Type | Description | Required |
|------|------|-------------|----------|
| enabled | boolean | Enabled determines if vLLM data should be persisted | false |
| size | string | Size is the size of the persistent volume | false |
| storageClass | string | StorageClass is the storage class to use for persistent volumes | false |
### LMDeployment.spec.vllm.resources

Resources defines the resource requirements for vLLM pods

| Name | Type | Description | Required |
|------|------|-------------|----------|
| limits | map[string]int or string | Limits describes the maximum amount of compute resources allowed | false |
| requests | map[string]int or string | Requests describes the minimum amount of compute resources required | false |
### LMDeployment.spec.vllm.service

Service defines the service configuration for vLLM

| Name | Type | Description | Required |
|------|------|-------------|----------|
| port | integer | Port is the port to expose the service<br/>*Format*: int32<br/>*Minimum*: 1<br/>*Maximum*: 65535<br/> | false |
| type | enum | Type is the type of service to expose<br/>*Enum*: ClusterIP, NodePort, LoadBalancer<br/> | false |
### LMDeployment.spec.vllm.volumeMounts[index]

VolumeMount describes a mounting of a Volume within a container.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| mountPath | string | Path within the container at which the volume should be mounted.  Must<br/>not contain ':'. | true |
| name | string | This must match the Name of a Volume. | true |
| mountPropagation | string | mountPropagation determines how mounts are propagated from the host<br/>to container and the other way around.<br/>When not set, MountPropagationNone is used.<br/>This field is beta in 1.10.<br/>When RecursiveReadOnly is set to IfPossible or to Enabled, MountPropagation must be None or unspecified<br/>(which defaults to None). | false |
| readOnly | boolean | Mounted read-only if true, read-write otherwise (false or unspecified).<br/>Defaults to false. | false |
| recursiveReadOnly | string | RecursiveReadOnly specifies whether read-only mounts should be handled<br/>recursively.<br/><br/>If ReadOnly is false, this field has no meaning and must be unspecified.<br/><br/>If ReadOnly is true, and this field is set to Disabled, the mount is not made<br/>recursively read-only.  If this field is set to IfPossible, the mount is made<br/>recursively read-only, if it is supported by the container runtime.  If this<br/>field is set to Enabled, the mount is made recursively read-only if it is<br/>supported by the container runtime, otherwise the pod will not be started and<br/>an error will be generated to indicate the reason.<br/><br/>If this field is set to IfPossible or Enabled, MountPropagation must be set to<br/>None (or be unspecified, which defaults to None).<br/><br/>If this field is not specified, it is treated as an equivalent of Disabled. | false |
| subPath | string | Path within the volume from which the container's volume should be mounted.<br/>Defaults to "" (volume's root). | false |
| subPathExpr | string | Expanded path within the volume from which the container's volume should be mounted.<br/>Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.<br/>Defaults to "" (volume's root).<br/>SubPathExpr and SubPath are mutually exclusive. | false |
### LMDeployment.spec.vllm.volumes[index]

Volume represents a named volume in a pod that may be accessed by any container in the pod.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | name of the volume.<br/>Must be a DNS_LABEL and unique within the pod.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names | true |
| [awsElasticBlockStore](#lmdeploymentspecvllmvolumesindexawselasticblockstore) | object | awsElasticBlockStore represents an AWS Disk resource that is attached to a<br/>kubelet's host machine and then exposed to the pod.<br/>Deprecated: AWSElasticBlockStore is deprecated. All operations for the in-tree<br/>awsElasticBlockStore type are redirected to the ebs.csi.aws.com CSI driver.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore | false |
| [azureDisk](#lmdeploymentspecvllmvolumesindexazuredisk) | object | azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.<br/>Deprecated: AzureDisk is deprecated. All operations for the in-tree azureDisk type<br/>are redirected to the disk.csi.azure.com CSI driver. | false |
| [azureFile](#lmdeploymentspecvllmvolumesindexazurefile) | object | azureFile represents an Azure File Service mount on the host and bind mount to the pod.<br/>Deprecated: AzureFile is deprecated. All operations for the in-tree azureFile type<br/>are redirected to the file.csi.azure.com CSI driver. | false |
| [cephfs](#lmdeploymentspecvllmvolumesindexcephfs) | object | cephFS represents a Ceph FS mount on the host that shares a pod's lifetime.<br/>Deprecated: CephFS is deprecated and the in-tree cephfs type is no longer supported. | false |
| [cinder](#lmdeploymentspecvllmvolumesindexcinder) | object | cinder represents a cinder volume attached and mounted on kubelets host machine.<br/>Deprecated: Cinder is deprecated. All operations for the in-tree cinder type<br/>are redirected to the cinder.csi.openstack.org CSI driver.<br/>More info: https://examples.k8s.io/mysql-cinder-pd/README.md | false |
| [configMap](#lmdeploymentspecvllmvolumesindexconfigmap) | object | configMap represents a configMap that should populate this volume | false |
| [csi](#lmdeploymentspecvllmvolumesindexcsi) | object | csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers. | false |
| [downwardAPI](#lmdeploymentspecvllmvolumesindexdownwardapi) | object | downwardAPI represents downward API about the pod that should populate this volume | false |
| [emptyDir](#lmdeploymentspecvllmvolumesindexemptydir) | object | emptyDir represents a temporary directory that shares a pod's lifetime.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir | false |
| [ephemeral](#lmdeploymentspecvllmvolumesindexephemeral) | object | ephemeral represents a volume that is handled by a cluster storage driver.<br/>The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts,<br/>and deleted when the pod is removed.<br/><br/>Use this if:<br/>a) the volume is only needed while the pod runs,<br/>b) features of normal volumes like restoring from snapshot or capacity<br/>   tracking are needed,<br/>c) the storage driver is specified through a storage class, and<br/>d) the storage driver supports dynamic volume provisioning through<br/>   a PersistentVolumeClaim (see EphemeralVolumeSource for more<br/>   information on the connection between this volume type<br/>   and PersistentVolumeClaim).<br/><br/>Use PersistentVolumeClaim or one of the vendor-specific<br/>APIs for volumes that persist for longer than the lifecycle<br/>of an individual pod.<br/><br/>Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to<br/>be used that way - see the documentation of the driver for<br/>more information.<br/><br/>A pod can use both types of ephemeral volumes and<br/>persistent volumes at the same time. | false |
| [fc](#lmdeploymentspecvllmvolumesindexfc) | object | fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod. | false |
| [flexVolume](#lmdeploymentspecvllmvolumesindexflexvolume) | object | flexVolume represents a generic volume resource that is<br/>provisioned/attached using an exec based plugin.<br/>Deprecated: FlexVolume is deprecated. Consider using a CSIDriver instead. | false |
| [flocker](#lmdeploymentspecvllmvolumesindexflocker) | object | flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running.<br/>Deprecated: Flocker is deprecated and the in-tree flocker type is no longer supported. | false |
| [gcePersistentDisk](#lmdeploymentspecvllmvolumesindexgcepersistentdisk) | object | gcePersistentDisk represents a GCE Disk resource that is attached to a<br/>kubelet's host machine and then exposed to the pod.<br/>Deprecated: GCEPersistentDisk is deprecated. All operations for the in-tree<br/>gcePersistentDisk type are redirected to the pd.csi.storage.gke.io CSI driver.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk | false |
| [gitRepo](#lmdeploymentspecvllmvolumesindexgitrepo) | object | gitRepo represents a git repository at a particular revision.<br/>Deprecated: GitRepo is deprecated. To provision a container with a git repo, mount an<br/>EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir<br/>into the Pod's container. | false |
| [glusterfs](#lmdeploymentspecvllmvolumesindexglusterfs) | object | glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.<br/>Deprecated: Glusterfs is deprecated and the in-tree glusterfs type is no longer supported.<br/>More info: https://examples.k8s.io/volumes/glusterfs/README.md | false |
| [hostPath](#lmdeploymentspecvllmvolumesindexhostpath) | object | hostPath represents a pre-existing file or directory on the host<br/>machine that is directly exposed to the container. This is generally<br/>used for system agents or other privileged things that are allowed<br/>to see the host machine. Most containers will NOT need this.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath | false |
| [image](#lmdeploymentspecvllmvolumesindeximage) | object | image represents an OCI object (a container image or artifact) pulled and mounted on the kubelet's host machine.<br/>The volume is resolved at pod startup depending on which PullPolicy value is provided:<br/><br/>- Always: the kubelet always attempts to pull the reference. Container creation will fail If the pull fails.<br/>- Never: the kubelet never pulls the reference and only uses a local image or artifact. Container creation will fail if the reference isn't present.<br/>- IfNotPresent: the kubelet pulls if the reference isn't already present on disk. Container creation will fail if the reference isn't present and the pull fails.<br/><br/>The volume gets re-resolved if the pod gets deleted and recreated, which means that new remote content will become available on pod recreation.<br/>A failure to resolve or pull the image during pod startup will block containers from starting and may add significant latency. Failures will be retried using normal volume backoff and will be reported on the pod reason and message.<br/>The types of objects that may be mounted by this volume are defined by the container runtime implementation on a host machine and at minimum must include all valid types supported by the container image field.<br/>The OCI object gets mounted in a single directory (spec.containers[*].volumeMounts.mountPath) by merging the manifest layers in the same way as for container images.<br/>The volume will be mounted read-only (ro) and non-executable files (noexec).<br/>Sub path mounts for containers are not supported (spec.containers[*].volumeMounts.subpath) before 1.33.<br/>The field spec.securityContext.fsGroupChangePolicy has no effect on this volume type. | false |
| [iscsi](#lmdeploymentspecvllmvolumesindexiscsi) | object | iscsi represents an ISCSI Disk resource that is attached to a<br/>kubelet's host machine and then exposed to the pod.<br/>More info: https://examples.k8s.io/volumes/iscsi/README.md | false |
| [nfs](#lmdeploymentspecvllmvolumesindexnfs) | object | nfs represents an NFS mount on the host that shares a pod's lifetime<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs | false |
| [persistentVolumeClaim](#lmdeploymentspecvllmvolumesindexpersistentvolumeclaim) | object | persistentVolumeClaimVolumeSource represents a reference to a<br/>PersistentVolumeClaim in the same namespace.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims | false |
| [photonPersistentDisk](#lmdeploymentspecvllmvolumesindexphotonpersistentdisk) | object | photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine.<br/>Deprecated: PhotonPersistentDisk is deprecated and the in-tree photonPersistentDisk type is no longer supported. | false |
| [portworxVolume](#lmdeploymentspecvllmvolumesindexportworxvolume) | object | portworxVolume represents a portworx volume attached and mounted on kubelets host machine.<br/>Deprecated: PortworxVolume is deprecated. All operations for the in-tree portworxVolume type<br/>are redirected to the pxd.portworx.com CSI driver when the CSIMigrationPortworx feature-gate<br/>is on. | false |
| [projected](#lmdeploymentspecvllmvolumesindexprojected) | object | projected items for all in one resources secrets, configmaps, and downward API | false |
| [quobyte](#lmdeploymentspecvllmvolumesindexquobyte) | object | quobyte represents a Quobyte mount on the host that shares a pod's lifetime.<br/>Deprecated: Quobyte is deprecated and the in-tree quobyte type is no longer supported. | false |
| [rbd](#lmdeploymentspecvllmvolumesindexrbd) | object | rbd represents a Rados Block Device mount on the host that shares a pod's lifetime.<br/>Deprecated: RBD is deprecated and the in-tree rbd type is no longer supported.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md | false |
| [scaleIO](#lmdeploymentspecvllmvolumesindexscaleio) | object | scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.<br/>Deprecated: ScaleIO is deprecated and the in-tree scaleIO type is no longer supported. | false |
| [secret](#lmdeploymentspecvllmvolumesindexsecret) | object | secret represents a secret that should populate this volume.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#secret | false |
| [storageos](#lmdeploymentspecvllmvolumesindexstorageos) | object | storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.<br/>Deprecated: StorageOS is deprecated and the in-tree storageos type is no longer supported. | false |
| [vsphereVolume](#lmdeploymentspecvllmvolumesindexvspherevolume) | object | vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine.<br/>Deprecated: VsphereVolume is deprecated. All operations for the in-tree vsphereVolume type<br/>are redirected to the csi.vsphere.vmware.com CSI driver. | false |
### LMDeployment.spec.vllm.volumes[index].awsElasticBlockStore

awsElasticBlockStore represents an AWS Disk resource that is attached to a<br/>kubelet's host machine and then exposed to the pod.<br/>Deprecated: AWSElasticBlockStore is deprecated. All operations for the in-tree<br/>awsElasticBlockStore type are redirected to the ebs.csi.aws.com CSI driver.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore

| Name | Type | Description | Required |
|------|------|-------------|----------|
| volumeID | string | volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume).<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore | true |
| fsType | string | fsType is the filesystem type of the volume that you want to mount.<br/>Tip: Ensure that the filesystem type is supported by the host operating system.<br/>Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore | false |
| partition | integer | partition is the partition in the volume that you want to mount.<br/>If omitted, the default is to mount by volume name.<br/>Examples: For volume /dev/sda1, you specify the partition as "1".<br/>Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).<br/>*Format*: int32<br/> | false |
| readOnly | boolean | readOnly value true will force the readOnly setting in VolumeMounts.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore | false |
### LMDeployment.spec.vllm.volumes[index].azureDisk

azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.<br/>Deprecated: AzureDisk is deprecated. All operations for the in-tree azureDisk type<br/>are redirected to the disk.csi.azure.com CSI driver.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| diskName | string | diskName is the Name of the data disk in the blob storage | true |
| diskURI | string | diskURI is the URI of data disk in the blob storage | true |
| cachingMode | string | cachingMode is the Host Caching mode: None, Read Only, Read Write. | false |
| fsType | string | fsType is Filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br/>*Default*: ext4<br/> | false |
| kind | string | kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared | false |
| readOnly | boolean | readOnly Defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts.<br/>*Default*: false<br/> | false |
### LMDeployment.spec.vllm.volumes[index].azureFile

azureFile represents an Azure File Service mount on the host and bind mount to the pod.<br/>Deprecated: AzureFile is deprecated. All operations for the in-tree azureFile type<br/>are redirected to the file.csi.azure.com CSI driver.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| secretName | string | secretName is the  name of secret that contains Azure Storage Account Name and Key | true |
| shareName | string | shareName is the azure share Name | true |
| readOnly | boolean | readOnly defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts. | false |
### LMDeployment.spec.vllm.volumes[index].cephfs

cephFS represents a Ceph FS mount on the host that shares a pod's lifetime.<br/>Deprecated: CephFS is deprecated and the in-tree cephfs type is no longer supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| monitors | []string | monitors is Required: Monitors is a collection of Ceph monitors<br/>More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | true |
| path | string | path is Optional: Used as the mounted root, rather than the full Ceph tree, default is / | false |
| readOnly | boolean | readOnly is Optional: Defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts.<br/>More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | false |
| secretFile | string | secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret<br/>More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | false |
| [secretRef](#lmdeploymentspecvllmvolumesindexcephfssecretref) | object | secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty.<br/>More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | false |
| user | string | user is optional: User is the rados user name, default is admin<br/>More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it | false |
### LMDeployment.spec.vllm.volumes[index].cephfs.secretRef

secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty.<br/>More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.vllm.volumes[index].cinder

cinder represents a cinder volume attached and mounted on kubelets host machine.<br/>Deprecated: Cinder is deprecated. All operations for the in-tree cinder type<br/>are redirected to the cinder.csi.openstack.org CSI driver.<br/>More info: https://examples.k8s.io/mysql-cinder-pd/README.md

| Name | Type | Description | Required |
|------|------|-------------|----------|
| volumeID | string | volumeID used to identify the volume in cinder.<br/>More info: https://examples.k8s.io/mysql-cinder-pd/README.md | true |
| fsType | string | fsType is the filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br/>More info: https://examples.k8s.io/mysql-cinder-pd/README.md | false |
| readOnly | boolean | readOnly defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts.<br/>More info: https://examples.k8s.io/mysql-cinder-pd/README.md | false |
| [secretRef](#lmdeploymentspecvllmvolumesindexcindersecretref) | object | secretRef is optional: points to a secret object containing parameters used to connect<br/>to OpenStack. | false |
### LMDeployment.spec.vllm.volumes[index].cinder.secretRef

secretRef is optional: points to a secret object containing parameters used to connect<br/>to OpenStack.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.vllm.volumes[index].configMap

configMap represents a configMap that should populate this volume

| Name | Type | Description | Required |
|------|------|-------------|----------|
| defaultMode | integer | defaultMode is optional: mode bits used to set permissions on created files by default.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>Defaults to 0644.<br/>Directories within the path are not affected by this setting.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
| [items](#lmdeploymentspecvllmvolumesindexconfigmapitemsindex) | []object | items if unspecified, each key-value pair in the Data field of the referenced<br/>ConfigMap will be projected into the volume as a file whose name is the<br/>key and content is the value. If specified, the listed keys will be<br/>projected into the specified paths, and unlisted keys will not be<br/>present. If a key is specified which is not present in the ConfigMap,<br/>the volume setup will error unless it is marked optional. Paths must be<br/>relative and may not contain the '..' path or start with '..'. | false |
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
| optional | boolean | optional specify whether the ConfigMap or its keys must be defined | false |
### LMDeployment.spec.vllm.volumes[index].configMap.items[index]

Maps a string key to a path within a volume.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the key to project. | true |
| path | string | path is the relative path of the file to map the key to.<br/>May not be an absolute path.<br/>May not contain the path element '..'.<br/>May not start with the string '..'. | true |
| mode | integer | mode is Optional: mode bits used to set permissions on this file.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>If not specified, the volume defaultMode will be used.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
### LMDeployment.spec.vllm.volumes[index].csi

csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| driver | string | driver is the name of the CSI driver that handles this volume.<br/>Consult with your admin for the correct name as registered in the cluster. | true |
| fsType | string | fsType to mount. Ex. "ext4", "xfs", "ntfs".<br/>If not provided, the empty value is passed to the associated CSI driver<br/>which will determine the default filesystem to apply. | false |
| [nodePublishSecretRef](#lmdeploymentspecvllmvolumesindexcsinodepublishsecretref) | object | nodePublishSecretRef is a reference to the secret object containing<br/>sensitive information to pass to the CSI driver to complete the CSI<br/>NodePublishVolume and NodeUnpublishVolume calls.<br/>This field is optional, and  may be empty if no secret is required. If the<br/>secret object contains more than one secret, all secret references are passed. | false |
| readOnly | boolean | readOnly specifies a read-only configuration for the volume.<br/>Defaults to false (read/write). | false |
| volumeAttributes | map[string]string | volumeAttributes stores driver-specific properties that are passed to the CSI<br/>driver. Consult your driver's documentation for supported values. | false |
### LMDeployment.spec.vllm.volumes[index].csi.nodePublishSecretRef

nodePublishSecretRef is a reference to the secret object containing<br/>sensitive information to pass to the CSI driver to complete the CSI<br/>NodePublishVolume and NodeUnpublishVolume calls.<br/>This field is optional, and  may be empty if no secret is required. If the<br/>secret object contains more than one secret, all secret references are passed.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.vllm.volumes[index].downwardAPI

downwardAPI represents downward API about the pod that should populate this volume

| Name | Type | Description | Required |
|------|------|-------------|----------|
| defaultMode | integer | Optional: mode bits to use on created files by default. Must be a<br/>Optional: mode bits used to set permissions on created files by default.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>Defaults to 0644.<br/>Directories within the path are not affected by this setting.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
| [items](#lmdeploymentspecvllmvolumesindexdownwardapiitemsindex) | []object | Items is a list of downward API volume file | false |
### LMDeployment.spec.vllm.volumes[index].downwardAPI.items[index]

DownwardAPIVolumeFile represents information to create the file containing the pod field

| Name | Type | Description | Required |
|------|------|-------------|----------|
| path | string | Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..' | true |
| [fieldRef](#lmdeploymentspecvllmvolumesindexdownwardapiitemsindexfieldref) | object | Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported. | false |
| mode | integer | Optional: mode bits used to set permissions on this file, must be an octal value<br/>between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>If not specified, the volume defaultMode will be used.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
| [resourceFieldRef](#lmdeploymentspecvllmvolumesindexdownwardapiitemsindexresourcefieldref) | object | Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported. | false |
### LMDeployment.spec.vllm.volumes[index].downwardAPI.items[index].fieldRef

Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| fieldPath | string | Path of the field to select in the specified API version. | true |
| apiVersion | string | Version of the schema the FieldPath is written in terms of, defaults to "v1". | false |
### LMDeployment.spec.vllm.volumes[index].downwardAPI.items[index].resourceFieldRef

Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| resource | string | Required: resource to select | true |
| containerName | string | Container name: required for volumes, optional for env vars | false |
| divisor | int or string | Specifies the output format of the exposed resources, defaults to "1" | false |
### LMDeployment.spec.vllm.volumes[index].emptyDir

emptyDir represents a temporary directory that shares a pod's lifetime.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir

| Name | Type | Description | Required |
|------|------|-------------|----------|
| medium | string | medium represents what type of storage medium should back this directory.<br/>The default is "" which means to use the node's default medium.<br/>Must be an empty string (default) or Memory.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir | false |
| sizeLimit | int or string | sizeLimit is the total amount of local storage required for this EmptyDir volume.<br/>The size limit is also applicable for memory medium.<br/>The maximum usage on memory medium EmptyDir would be the minimum value between<br/>the SizeLimit specified here and the sum of memory limits of all containers in a pod.<br/>The default is nil which means that the limit is undefined.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir | false |
### LMDeployment.spec.vllm.volumes[index].ephemeral

ephemeral represents a volume that is handled by a cluster storage driver.<br/>The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts,<br/>and deleted when the pod is removed.<br/><br/>Use this if:<br/>a) the volume is only needed while the pod runs,<br/>b) features of normal volumes like restoring from snapshot or capacity<br/>   tracking are needed,<br/>c) the storage driver is specified through a storage class, and<br/>d) the storage driver supports dynamic volume provisioning through<br/>   a PersistentVolumeClaim (see EphemeralVolumeSource for more<br/>   information on the connection between this volume type<br/>   and PersistentVolumeClaim).<br/><br/>Use PersistentVolumeClaim or one of the vendor-specific<br/>APIs for volumes that persist for longer than the lifecycle<br/>of an individual pod.<br/><br/>Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to<br/>be used that way - see the documentation of the driver for<br/>more information.<br/><br/>A pod can use both types of ephemeral volumes and<br/>persistent volumes at the same time.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [volumeClaimTemplate](#lmdeploymentspecvllmvolumesindexephemeralvolumeclaimtemplate) | object | Will be used to create a stand-alone PVC to provision the volume.<br/>The pod in which this EphemeralVolumeSource is embedded will be the<br/>owner of the PVC, i.e. the PVC will be deleted together with the<br/>pod.  The name of the PVC will be `&lt;pod name&gt;-&lt;volume name&gt;` where<br/>`&lt;volume name&gt;` is the name from the `PodSpec.Volumes` array<br/>entry. Pod validation will reject the pod if the concatenated name<br/>is not valid for a PVC (for example, too long).<br/><br/>An existing PVC with that name that is not owned by the pod<br/>will *not* be used for the pod to avoid using an unrelated<br/>volume by mistake. Starting the pod is then blocked until<br/>the unrelated PVC is removed. If such a pre-created PVC is<br/>meant to be used by the pod, the PVC has to updated with an<br/>owner reference to the pod once the pod exists. Normally<br/>this should not be necessary, but it may be useful when<br/>manually reconstructing a broken cluster.<br/><br/>This field is read-only and no changes will be made by Kubernetes<br/>to the PVC after it has been created.<br/><br/>Required, must not be nil. | false |
### LMDeployment.spec.vllm.volumes[index].ephemeral.volumeClaimTemplate

Will be used to create a stand-alone PVC to provision the volume.<br/>The pod in which this EphemeralVolumeSource is embedded will be the<br/>owner of the PVC, i.e. the PVC will be deleted together with the<br/>pod.  The name of the PVC will be `&lt;pod name&gt;-&lt;volume name&gt;` where<br/>`&lt;volume name&gt;` is the name from the `PodSpec.Volumes` array<br/>entry. Pod validation will reject the pod if the concatenated name<br/>is not valid for a PVC (for example, too long).<br/><br/>An existing PVC with that name that is not owned by the pod<br/>will *not* be used for the pod to avoid using an unrelated<br/>volume by mistake. Starting the pod is then blocked until<br/>the unrelated PVC is removed. If such a pre-created PVC is<br/>meant to be used by the pod, the PVC has to updated with an<br/>owner reference to the pod once the pod exists. Normally<br/>this should not be necessary, but it may be useful when<br/>manually reconstructing a broken cluster.<br/><br/>This field is read-only and no changes will be made by Kubernetes<br/>to the PVC after it has been created.<br/><br/>Required, must not be nil.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [spec](#lmdeploymentspecvllmvolumesindexephemeralvolumeclaimtemplatespec) | object | The specification for the PersistentVolumeClaim. The entire content is<br/>copied unchanged into the PVC that gets created from this<br/>template. The same fields as in a PersistentVolumeClaim<br/>are also valid here. | true |
| metadata | object | May contain labels and annotations that will be copied into the PVC<br/>when creating it. No other fields are allowed and will be rejected during<br/>validation. | false |
### LMDeployment.spec.vllm.volumes[index].ephemeral.volumeClaimTemplate.spec

The specification for the PersistentVolumeClaim. The entire content is<br/>copied unchanged into the PVC that gets created from this<br/>template. The same fields as in a PersistentVolumeClaim<br/>are also valid here.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| accessModes | []string | accessModes contains the desired access modes the volume should have.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 | false |
| [dataSource](#lmdeploymentspecvllmvolumesindexephemeralvolumeclaimtemplatespecdatasource) | object | dataSource field can be used to specify either:<br/>* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)<br/>* An existing PVC (PersistentVolumeClaim)<br/>If the provisioner or an external controller can support the specified data source,<br/>it will create a new volume based on the contents of the specified data source.<br/>When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,<br/>and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.<br/>If the namespace is specified, then dataSourceRef will not be copied to dataSource. | false |
| [dataSourceRef](#lmdeploymentspecvllmvolumesindexephemeralvolumeclaimtemplatespecdatasourceref) | object | dataSourceRef specifies the object from which to populate the volume with data, if a non-empty<br/>volume is desired. This may be any object from a non-empty API group (non<br/>core object) or a PersistentVolumeClaim object.<br/>When this field is specified, volume binding will only succeed if the type of<br/>the specified object matches some installed volume populator or dynamic<br/>provisioner.<br/>This field will replace the functionality of the dataSource field and as such<br/>if both fields are non-empty, they must have the same value. For backwards<br/>compatibility, when namespace isn't specified in dataSourceRef,<br/>both fields (dataSource and dataSourceRef) will be set to the same<br/>value automatically if one of them is empty and the other is non-empty.<br/>When namespace is specified in dataSourceRef,<br/>dataSource isn't set to the same value and must be empty.<br/>There are three important differences between dataSource and dataSourceRef:<br/>* While dataSource only allows two specific types of objects, dataSourceRef<br/>  allows any non-core object, as well as PersistentVolumeClaim objects.<br/>* While dataSource ignores disallowed values (dropping them), dataSourceRef<br/>  preserves all values, and generates an error if a disallowed value is<br/>  specified.<br/>* While dataSource only allows local objects, dataSourceRef allows objects<br/>  in any namespaces.<br/>(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.<br/>(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled. | false |
| [resources](#lmdeploymentspecvllmvolumesindexephemeralvolumeclaimtemplatespecresources) | object | resources represents the minimum resources the volume should have.<br/>If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements<br/>that are lower than previous value but must still be higher than capacity recorded in the<br/>status field of the claim.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources | false |
| [selector](#lmdeploymentspecvllmvolumesindexephemeralvolumeclaimtemplatespecselector) | object | selector is a label query over volumes to consider for binding. | false |
| storageClassName | string | storageClassName is the name of the StorageClass required by the claim.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1 | false |
| volumeAttributesClassName | string | volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.<br/>If specified, the CSI driver will create or update the volume with the attributes defined<br/>in the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,<br/>it can be changed after the claim is created. An empty string value means that no VolumeAttributesClass<br/>will be applied to the claim but it's not allowed to reset this field to empty string once it is set.<br/>If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClass<br/>will be set by the persistentvolume controller if it exists.<br/>If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will be<br/>set to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resource<br/>exists.<br/>More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/<br/>(Beta) Using this field requires the VolumeAttributesClass feature gate to be enabled (off by default). | false |
| volumeMode | string | volumeMode defines what type of volume is required by the claim.<br/>Value of Filesystem is implied when not included in claim spec. | false |
| volumeName | string | volumeName is the binding reference to the PersistentVolume backing this claim. | false |
### LMDeployment.spec.vllm.volumes[index].ephemeral.volumeClaimTemplate.spec.dataSource

dataSource field can be used to specify either:<br/>* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)<br/>* An existing PVC (PersistentVolumeClaim)<br/>If the provisioner or an external controller can support the specified data source,<br/>it will create a new volume based on the contents of the specified data source.<br/>When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,<br/>and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.<br/>If the namespace is specified, then dataSourceRef will not be copied to dataSource.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| kind | string | Kind is the type of resource being referenced | true |
| name | string | Name is the name of resource being referenced | true |
| apiGroup | string | APIGroup is the group for the resource being referenced.<br/>If APIGroup is not specified, the specified Kind must be in the core API group.<br/>For any other third-party types, APIGroup is required. | false |
### LMDeployment.spec.vllm.volumes[index].ephemeral.volumeClaimTemplate.spec.dataSourceRef

dataSourceRef specifies the object from which to populate the volume with data, if a non-empty<br/>volume is desired. This may be any object from a non-empty API group (non<br/>core object) or a PersistentVolumeClaim object.<br/>When this field is specified, volume binding will only succeed if the type of<br/>the specified object matches some installed volume populator or dynamic<br/>provisioner.<br/>This field will replace the functionality of the dataSource field and as such<br/>if both fields are non-empty, they must have the same value. For backwards<br/>compatibility, when namespace isn't specified in dataSourceRef,<br/>both fields (dataSource and dataSourceRef) will be set to the same<br/>value automatically if one of them is empty and the other is non-empty.<br/>When namespace is specified in dataSourceRef,<br/>dataSource isn't set to the same value and must be empty.<br/>There are three important differences between dataSource and dataSourceRef:<br/>* While dataSource only allows two specific types of objects, dataSourceRef<br/>  allows any non-core object, as well as PersistentVolumeClaim objects.<br/>* While dataSource ignores disallowed values (dropping them), dataSourceRef<br/>  preserves all values, and generates an error if a disallowed value is<br/>  specified.<br/>* While dataSource only allows local objects, dataSourceRef allows objects<br/>  in any namespaces.<br/>(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.<br/>(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| kind | string | Kind is the type of resource being referenced | true |
| name | string | Name is the name of resource being referenced | true |
| apiGroup | string | APIGroup is the group for the resource being referenced.<br/>If APIGroup is not specified, the specified Kind must be in the core API group.<br/>For any other third-party types, APIGroup is required. | false |
| namespace | string | Namespace is the namespace of resource being referenced<br/>Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.<br/>(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled. | false |
### LMDeployment.spec.vllm.volumes[index].ephemeral.volumeClaimTemplate.spec.resources

resources represents the minimum resources the volume should have.<br/>If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements<br/>that are lower than previous value but must still be higher than capacity recorded in the<br/>status field of the claim.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources

| Name | Type | Description | Required |
|------|------|-------------|----------|
| limits | map[string]int or string | Limits describes the maximum amount of compute resources allowed.<br/>More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ | false |
| requests | map[string]int or string | Requests describes the minimum amount of compute resources required.<br/>If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,<br/>otherwise to an implementation-defined value. Requests cannot exceed Limits.<br/>More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ | false |
### LMDeployment.spec.vllm.volumes[index].ephemeral.volumeClaimTemplate.spec.selector

selector is a label query over volumes to consider for binding.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecvllmvolumesindexephemeralvolumeclaimtemplatespecselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.vllm.volumes[index].ephemeral.volumeClaimTemplate.spec.selector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.vllm.volumes[index].fc

fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| fsType | string | fsType is the filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. | false |
| lun | integer | lun is Optional: FC target lun number<br/>*Format*: int32<br/> | false |
| readOnly | boolean | readOnly is Optional: Defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts. | false |
| targetWWNs | []string | targetWWNs is Optional: FC target worldwide names (WWNs) | false |
| wwids | []string | wwids Optional: FC volume world wide identifiers (wwids)<br/>Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously. | false |
### LMDeployment.spec.vllm.volumes[index].flexVolume

flexVolume represents a generic volume resource that is<br/>provisioned/attached using an exec based plugin.<br/>Deprecated: FlexVolume is deprecated. Consider using a CSIDriver instead.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| driver | string | driver is the name of the driver to use for this volume. | true |
| fsType | string | fsType is the filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script. | false |
| options | map[string]string | options is Optional: this field holds extra command options if any. | false |
| readOnly | boolean | readOnly is Optional: defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts. | false |
| [secretRef](#lmdeploymentspecvllmvolumesindexflexvolumesecretref) | object | secretRef is Optional: secretRef is reference to the secret object containing<br/>sensitive information to pass to the plugin scripts. This may be<br/>empty if no secret object is specified. If the secret object<br/>contains more than one secret, all secrets are passed to the plugin<br/>scripts. | false |
### LMDeployment.spec.vllm.volumes[index].flexVolume.secretRef

secretRef is Optional: secretRef is reference to the secret object containing<br/>sensitive information to pass to the plugin scripts. This may be<br/>empty if no secret object is specified. If the secret object<br/>contains more than one secret, all secrets are passed to the plugin<br/>scripts.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.vllm.volumes[index].flocker

flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running.<br/>Deprecated: Flocker is deprecated and the in-tree flocker type is no longer supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| datasetName | string | datasetName is Name of the dataset stored as metadata -&gt; name on the dataset for Flocker<br/>should be considered as deprecated | false |
| datasetUUID | string | datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset | false |
### LMDeployment.spec.vllm.volumes[index].gcePersistentDisk

gcePersistentDisk represents a GCE Disk resource that is attached to a<br/>kubelet's host machine and then exposed to the pod.<br/>Deprecated: GCEPersistentDisk is deprecated. All operations for the in-tree<br/>gcePersistentDisk type are redirected to the pd.csi.storage.gke.io CSI driver.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk

| Name | Type | Description | Required |
|------|------|-------------|----------|
| pdName | string | pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk | true |
| fsType | string | fsType is filesystem type of the volume that you want to mount.<br/>Tip: Ensure that the filesystem type is supported by the host operating system.<br/>Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk | false |
| partition | integer | partition is the partition in the volume that you want to mount.<br/>If omitted, the default is to mount by volume name.<br/>Examples: For volume /dev/sda1, you specify the partition as "1".<br/>Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk<br/>*Format*: int32<br/> | false |
| readOnly | boolean | readOnly here will force the ReadOnly setting in VolumeMounts.<br/>Defaults to false.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk | false |
### LMDeployment.spec.vllm.volumes[index].gitRepo

gitRepo represents a git repository at a particular revision.<br/>Deprecated: GitRepo is deprecated. To provision a container with a git repo, mount an<br/>EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir<br/>into the Pod's container.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| repository | string | repository is the URL | true |
| directory | string | directory is the target directory name.<br/>Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the<br/>git repository.  Otherwise, if specified, the volume will contain the git repository in<br/>the subdirectory with the given name. | false |
| revision | string | revision is the commit hash for the specified revision. | false |
### LMDeployment.spec.vllm.volumes[index].glusterfs

glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.<br/>Deprecated: Glusterfs is deprecated and the in-tree glusterfs type is no longer supported.<br/>More info: https://examples.k8s.io/volumes/glusterfs/README.md

| Name | Type | Description | Required |
|------|------|-------------|----------|
| endpoints | string | endpoints is the endpoint name that details Glusterfs topology.<br/>More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod | true |
| path | string | path is the Glusterfs volume path.<br/>More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod | true |
| readOnly | boolean | readOnly here will force the Glusterfs volume to be mounted with read-only permissions.<br/>Defaults to false.<br/>More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod | false |
### LMDeployment.spec.vllm.volumes[index].hostPath

hostPath represents a pre-existing file or directory on the host<br/>machine that is directly exposed to the container. This is generally<br/>used for system agents or other privileged things that are allowed<br/>to see the host machine. Most containers will NOT need this.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath

| Name | Type | Description | Required |
|------|------|-------------|----------|
| path | string | path of the directory on the host.<br/>If the path is a symlink, it will follow the link to the real path.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath | true |
| type | string | type for HostPath Volume<br/>Defaults to ""<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath | false |
### LMDeployment.spec.vllm.volumes[index].image

image represents an OCI object (a container image or artifact) pulled and mounted on the kubelet's host machine.<br/>The volume is resolved at pod startup depending on which PullPolicy value is provided:<br/><br/>- Always: the kubelet always attempts to pull the reference. Container creation will fail If the pull fails.<br/>- Never: the kubelet never pulls the reference and only uses a local image or artifact. Container creation will fail if the reference isn't present.<br/>- IfNotPresent: the kubelet pulls if the reference isn't already present on disk. Container creation will fail if the reference isn't present and the pull fails.<br/><br/>The volume gets re-resolved if the pod gets deleted and recreated, which means that new remote content will become available on pod recreation.<br/>A failure to resolve or pull the image during pod startup will block containers from starting and may add significant latency. Failures will be retried using normal volume backoff and will be reported on the pod reason and message.<br/>The types of objects that may be mounted by this volume are defined by the container runtime implementation on a host machine and at minimum must include all valid types supported by the container image field.<br/>The OCI object gets mounted in a single directory (spec.containers[*].volumeMounts.mountPath) by merging the manifest layers in the same way as for container images.<br/>The volume will be mounted read-only (ro) and non-executable files (noexec).<br/>Sub path mounts for containers are not supported (spec.containers[*].volumeMounts.subpath) before 1.33.<br/>The field spec.securityContext.fsGroupChangePolicy has no effect on this volume type.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| pullPolicy | string | Policy for pulling OCI objects. Possible values are:<br/>Always: the kubelet always attempts to pull the reference. Container creation will fail If the pull fails.<br/>Never: the kubelet never pulls the reference and only uses a local image or artifact. Container creation will fail if the reference isn't present.<br/>IfNotPresent: the kubelet pulls if the reference isn't already present on disk. Container creation will fail if the reference isn't present and the pull fails.<br/>Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. | false |
| reference | string | Required: Image or artifact reference to be used.<br/>Behaves in the same way as pod.spec.containers[*].image.<br/>Pull secrets will be assembled in the same way as for the container image by looking up node credentials, SA image pull secrets, and pod spec image pull secrets.<br/>More info: https://kubernetes.io/docs/concepts/containers/images<br/>This field is optional to allow higher level config management to default or override<br/>container images in workload controllers like Deployments and StatefulSets. | false |
### LMDeployment.spec.vllm.volumes[index].iscsi

iscsi represents an ISCSI Disk resource that is attached to a<br/>kubelet's host machine and then exposed to the pod.<br/>More info: https://examples.k8s.io/volumes/iscsi/README.md

| Name | Type | Description | Required |
|------|------|-------------|----------|
| iqn | string | iqn is the target iSCSI Qualified Name. | true |
| lun | integer | lun represents iSCSI Target Lun number.<br/>*Format*: int32<br/> | true |
| targetPortal | string | targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port<br/>is other than default (typically TCP ports 860 and 3260). | true |
| chapAuthDiscovery | boolean | chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication | false |
| chapAuthSession | boolean | chapAuthSession defines whether support iSCSI Session CHAP authentication | false |
| fsType | string | fsType is the filesystem type of the volume that you want to mount.<br/>Tip: Ensure that the filesystem type is supported by the host operating system.<br/>Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi | false |
| initiatorName | string | initiatorName is the custom iSCSI Initiator Name.<br/>If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface<br/>&lt;target portal&gt;:&lt;volume name&gt; will be created for the connection. | false |
| iscsiInterface | string | iscsiInterface is the interface Name that uses an iSCSI transport.<br/>Defaults to 'default' (tcp).<br/>*Default*: default<br/> | false |
| portals | []string | portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port<br/>is other than default (typically TCP ports 860 and 3260). | false |
| readOnly | boolean | readOnly here will force the ReadOnly setting in VolumeMounts.<br/>Defaults to false. | false |
| [secretRef](#lmdeploymentspecvllmvolumesindexiscsisecretref) | object | secretRef is the CHAP Secret for iSCSI target and initiator authentication | false |
### LMDeployment.spec.vllm.volumes[index].iscsi.secretRef

secretRef is the CHAP Secret for iSCSI target and initiator authentication

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.vllm.volumes[index].nfs

nfs represents an NFS mount on the host that shares a pod's lifetime<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs

| Name | Type | Description | Required |
|------|------|-------------|----------|
| path | string | path that is exported by the NFS server.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs | true |
| server | string | server is the hostname or IP address of the NFS server.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs | true |
| readOnly | boolean | readOnly here will force the NFS export to be mounted with read-only permissions.<br/>Defaults to false.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs | false |
### LMDeployment.spec.vllm.volumes[index].persistentVolumeClaim

persistentVolumeClaimVolumeSource represents a reference to a<br/>PersistentVolumeClaim in the same namespace.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims

| Name | Type | Description | Required |
|------|------|-------------|----------|
| claimName | string | claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.<br/>More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims | true |
| readOnly | boolean | readOnly Will force the ReadOnly setting in VolumeMounts.<br/>Default false. | false |
### LMDeployment.spec.vllm.volumes[index].photonPersistentDisk

photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine.<br/>Deprecated: PhotonPersistentDisk is deprecated and the in-tree photonPersistentDisk type is no longer supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| pdID | string | pdID is the ID that identifies Photon Controller persistent disk | true |
| fsType | string | fsType is the filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. | false |
### LMDeployment.spec.vllm.volumes[index].portworxVolume

portworxVolume represents a portworx volume attached and mounted on kubelets host machine.<br/>Deprecated: PortworxVolume is deprecated. All operations for the in-tree portworxVolume type<br/>are redirected to the pxd.portworx.com CSI driver when the CSIMigrationPortworx feature-gate<br/>is on.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| volumeID | string | volumeID uniquely identifies a Portworx volume | true |
| fsType | string | fSType represents the filesystem type to mount<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified. | false |
| readOnly | boolean | readOnly defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts. | false |
### LMDeployment.spec.vllm.volumes[index].projected

projected items for all in one resources secrets, configmaps, and downward API

| Name | Type | Description | Required |
|------|------|-------------|----------|
| defaultMode | integer | defaultMode are the mode bits used to set permissions on created files by default.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>Directories within the path are not affected by this setting.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
| [sources](#lmdeploymentspecvllmvolumesindexprojectedsourcesindex) | []object | sources is the list of volume projections. Each entry in this list<br/>handles one source. | false |
### LMDeployment.spec.vllm.volumes[index].projected.sources[index]

Projection that may be projected along with other supported volume types.<br/>Exactly one of these fields must be set.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [clusterTrustBundle](#lmdeploymentspecvllmvolumesindexprojectedsourcesindexclustertrustbundle) | object | ClusterTrustBundle allows a pod to access the `.spec.trustBundle` field<br/>of ClusterTrustBundle objects in an auto-updating file.<br/><br/>Alpha, gated by the ClusterTrustBundleProjection feature gate.<br/><br/>ClusterTrustBundle objects can either be selected by name, or by the<br/>combination of signer name and a label selector.<br/><br/>Kubelet performs aggressive normalization of the PEM contents written<br/>into the pod filesystem.  Esoteric PEM features such as inter-block<br/>comments and block headers are stripped.  Certificates are deduplicated.<br/>The ordering of certificates within the file is arbitrary, and Kubelet<br/>may change the order over time. | false |
| [configMap](#lmdeploymentspecvllmvolumesindexprojectedsourcesindexconfigmap) | object | configMap information about the configMap data to project | false |
| [downwardAPI](#lmdeploymentspecvllmvolumesindexprojectedsourcesindexdownwardapi) | object | downwardAPI information about the downwardAPI data to project | false |
| [secret](#lmdeploymentspecvllmvolumesindexprojectedsourcesindexsecret) | object | secret information about the secret data to project | false |
| [serviceAccountToken](#lmdeploymentspecvllmvolumesindexprojectedsourcesindexserviceaccounttoken) | object | serviceAccountToken is information about the serviceAccountToken data to project | false |
### LMDeployment.spec.vllm.volumes[index].projected.sources[index].clusterTrustBundle

ClusterTrustBundle allows a pod to access the `.spec.trustBundle` field<br/>of ClusterTrustBundle objects in an auto-updating file.<br/><br/>Alpha, gated by the ClusterTrustBundleProjection feature gate.<br/><br/>ClusterTrustBundle objects can either be selected by name, or by the<br/>combination of signer name and a label selector.<br/><br/>Kubelet performs aggressive normalization of the PEM contents written<br/>into the pod filesystem.  Esoteric PEM features such as inter-block<br/>comments and block headers are stripped.  Certificates are deduplicated.<br/>The ordering of certificates within the file is arbitrary, and Kubelet<br/>may change the order over time.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| path | string | Relative path from the volume root to write the bundle. | true |
| [labelSelector](#lmdeploymentspecvllmvolumesindexprojectedsourcesindexclustertrustbundlelabelselector) | object | Select all ClusterTrustBundles that match this label selector.  Only has<br/>effect if signerName is set.  Mutually-exclusive with name.  If unset,<br/>interpreted as "match nothing".  If set but empty, interpreted as "match<br/>everything". | false |
| name | string | Select a single ClusterTrustBundle by object name.  Mutually-exclusive<br/>with signerName and labelSelector. | false |
| optional | boolean | If true, don't block pod startup if the referenced ClusterTrustBundle(s)<br/>aren't available.  If using name, then the named ClusterTrustBundle is<br/>allowed not to exist.  If using signerName, then the combination of<br/>signerName and labelSelector is allowed to match zero<br/>ClusterTrustBundles. | false |
| signerName | string | Select all ClusterTrustBundles that match this signer name.<br/>Mutually-exclusive with name.  The contents of all selected<br/>ClusterTrustBundles will be unified and deduplicated. | false |
### LMDeployment.spec.vllm.volumes[index].projected.sources[index].clusterTrustBundle.labelSelector

Select all ClusterTrustBundles that match this label selector.  Only has<br/>effect if signerName is set.  Mutually-exclusive with name.  If unset,<br/>interpreted as "match nothing".  If set but empty, interpreted as "match<br/>everything".

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [matchExpressions](#lmdeploymentspecvllmvolumesindexprojectedsourcesindexclustertrustbundlelabelselectormatchexpressionsindex) | []object | matchExpressions is a list of label selector requirements. The requirements are ANDed. | false |
| matchLabels | map[string]string | matchLabels is a map of &#123;key,value&#125; pairs. A single &#123;key,value&#125; in the matchLabels<br/>map is equivalent to an element of matchExpressions, whose key field is "key", the<br/>operator is "In", and the values array contains only "value". The requirements are ANDed. | false |
### LMDeployment.spec.vllm.volumes[index].projected.sources[index].clusterTrustBundle.labelSelector.matchExpressions[index]

A label selector requirement is a selector that contains values, a key, and an operator that<br/>relates the key and values.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the label key that the selector applies to. | true |
| operator | string | operator represents a key's relationship to a set of values.<br/>Valid operators are In, NotIn, Exists and DoesNotExist. | true |
| values | []string | values is an array of string values. If the operator is In or NotIn,<br/>the values array must be non-empty. If the operator is Exists or DoesNotExist,<br/>the values array must be empty. This array is replaced during a strategic<br/>merge patch. | false |
### LMDeployment.spec.vllm.volumes[index].projected.sources[index].configMap

configMap information about the configMap data to project

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [items](#lmdeploymentspecvllmvolumesindexprojectedsourcesindexconfigmapitemsindex) | []object | items if unspecified, each key-value pair in the Data field of the referenced<br/>ConfigMap will be projected into the volume as a file whose name is the<br/>key and content is the value. If specified, the listed keys will be<br/>projected into the specified paths, and unlisted keys will not be<br/>present. If a key is specified which is not present in the ConfigMap,<br/>the volume setup will error unless it is marked optional. Paths must be<br/>relative and may not contain the '..' path or start with '..'. | false |
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
| optional | boolean | optional specify whether the ConfigMap or its keys must be defined | false |
### LMDeployment.spec.vllm.volumes[index].projected.sources[index].configMap.items[index]

Maps a string key to a path within a volume.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the key to project. | true |
| path | string | path is the relative path of the file to map the key to.<br/>May not be an absolute path.<br/>May not contain the path element '..'.<br/>May not start with the string '..'. | true |
| mode | integer | mode is Optional: mode bits used to set permissions on this file.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>If not specified, the volume defaultMode will be used.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
### LMDeployment.spec.vllm.volumes[index].projected.sources[index].downwardAPI

downwardAPI information about the downwardAPI data to project

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [items](#lmdeploymentspecvllmvolumesindexprojectedsourcesindexdownwardapiitemsindex) | []object | Items is a list of DownwardAPIVolume file | false |
### LMDeployment.spec.vllm.volumes[index].projected.sources[index].downwardAPI.items[index]

DownwardAPIVolumeFile represents information to create the file containing the pod field

| Name | Type | Description | Required |
|------|------|-------------|----------|
| path | string | Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..' | true |
| [fieldRef](#lmdeploymentspecvllmvolumesindexprojectedsourcesindexdownwardapiitemsindexfieldref) | object | Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported. | false |
| mode | integer | Optional: mode bits used to set permissions on this file, must be an octal value<br/>between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>If not specified, the volume defaultMode will be used.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
| [resourceFieldRef](#lmdeploymentspecvllmvolumesindexprojectedsourcesindexdownwardapiitemsindexresourcefieldref) | object | Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported. | false |
### LMDeployment.spec.vllm.volumes[index].projected.sources[index].downwardAPI.items[index].fieldRef

Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| fieldPath | string | Path of the field to select in the specified API version. | true |
| apiVersion | string | Version of the schema the FieldPath is written in terms of, defaults to "v1". | false |
### LMDeployment.spec.vllm.volumes[index].projected.sources[index].downwardAPI.items[index].resourceFieldRef

Selects a resource of the container: only resources limits and requests<br/>(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| resource | string | Required: resource to select | true |
| containerName | string | Container name: required for volumes, optional for env vars | false |
| divisor | int or string | Specifies the output format of the exposed resources, defaults to "1" | false |
### LMDeployment.spec.vllm.volumes[index].projected.sources[index].secret

secret information about the secret data to project

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [items](#lmdeploymentspecvllmvolumesindexprojectedsourcesindexsecretitemsindex) | []object | items if unspecified, each key-value pair in the Data field of the referenced<br/>Secret will be projected into the volume as a file whose name is the<br/>key and content is the value. If specified, the listed keys will be<br/>projected into the specified paths, and unlisted keys will not be<br/>present. If a key is specified which is not present in the Secret,<br/>the volume setup will error unless it is marked optional. Paths must be<br/>relative and may not contain the '..' path or start with '..'. | false |
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
| optional | boolean | optional field specify whether the Secret or its key must be defined | false |
### LMDeployment.spec.vllm.volumes[index].projected.sources[index].secret.items[index]

Maps a string key to a path within a volume.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the key to project. | true |
| path | string | path is the relative path of the file to map the key to.<br/>May not be an absolute path.<br/>May not contain the path element '..'.<br/>May not start with the string '..'. | true |
| mode | integer | mode is Optional: mode bits used to set permissions on this file.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>If not specified, the volume defaultMode will be used.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
### LMDeployment.spec.vllm.volumes[index].projected.sources[index].serviceAccountToken

serviceAccountToken is information about the serviceAccountToken data to project

| Name | Type | Description | Required |
|------|------|-------------|----------|
| path | string | path is the path relative to the mount point of the file to project the<br/>token into. | true |
| audience | string | audience is the intended audience of the token. A recipient of a token<br/>must identify itself with an identifier specified in the audience of the<br/>token, and otherwise should reject the token. The audience defaults to the<br/>identifier of the apiserver. | false |
| expirationSeconds | integer | expirationSeconds is the requested duration of validity of the service<br/>account token. As the token approaches expiration, the kubelet volume<br/>plugin will proactively rotate the service account token. The kubelet will<br/>start trying to rotate the token if the token is older than 80 percent of<br/>its time to live or if the token is older than 24 hours.Defaults to 1 hour<br/>and must be at least 10 minutes.<br/>*Format*: int64<br/> | false |
### LMDeployment.spec.vllm.volumes[index].quobyte

quobyte represents a Quobyte mount on the host that shares a pod's lifetime.<br/>Deprecated: Quobyte is deprecated and the in-tree quobyte type is no longer supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| registry | string | registry represents a single or multiple Quobyte Registry services<br/>specified as a string as host:port pair (multiple entries are separated with commas)<br/>which acts as the central registry for volumes | true |
| volume | string | volume is a string that references an already created Quobyte volume by name. | true |
| group | string | group to map volume access to<br/>Default is no group | false |
| readOnly | boolean | readOnly here will force the Quobyte volume to be mounted with read-only permissions.<br/>Defaults to false. | false |
| tenant | string | tenant owning the given Quobyte volume in the Backend<br/>Used with dynamically provisioned Quobyte volumes, value is set by the plugin | false |
| user | string | user to map volume access to<br/>Defaults to serivceaccount user | false |
### LMDeployment.spec.vllm.volumes[index].rbd

rbd represents a Rados Block Device mount on the host that shares a pod's lifetime.<br/>Deprecated: RBD is deprecated and the in-tree rbd type is no longer supported.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md

| Name | Type | Description | Required |
|------|------|-------------|----------|
| image | string | image is the rados image name.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | true |
| monitors | []string | monitors is a collection of Ceph monitors.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | true |
| fsType | string | fsType is the filesystem type of the volume that you want to mount.<br/>Tip: Ensure that the filesystem type is supported by the host operating system.<br/>Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd | false |
| keyring | string | keyring is the path to key ring for RBDUser.<br/>Default is /etc/ceph/keyring.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it<br/>*Default*: /etc/ceph/keyring<br/> | false |
| pool | string | pool is the rados pool name.<br/>Default is rbd.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it<br/>*Default*: rbd<br/> | false |
| readOnly | boolean | readOnly here will force the ReadOnly setting in VolumeMounts.<br/>Defaults to false.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | false |
| [secretRef](#lmdeploymentspecvllmvolumesindexrbdsecretref) | object | secretRef is name of the authentication secret for RBDUser. If provided<br/>overrides keyring.<br/>Default is nil.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it | false |
| user | string | user is the rados user name.<br/>Default is admin.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it<br/>*Default*: admin<br/> | false |
### LMDeployment.spec.vllm.volumes[index].rbd.secretRef

secretRef is name of the authentication secret for RBDUser. If provided<br/>overrides keyring.<br/>Default is nil.<br/>More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.vllm.volumes[index].scaleIO

scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.<br/>Deprecated: ScaleIO is deprecated and the in-tree scaleIO type is no longer supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| gateway | string | gateway is the host address of the ScaleIO API Gateway. | true |
| [secretRef](#lmdeploymentspecvllmvolumesindexscaleiosecretref) | object | secretRef references to the secret for ScaleIO user and other<br/>sensitive information. If this is not provided, Login operation will fail. | true |
| system | string | system is the name of the storage system as configured in ScaleIO. | true |
| fsType | string | fsType is the filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs".<br/>Default is "xfs".<br/>*Default*: xfs<br/> | false |
| protectionDomain | string | protectionDomain is the name of the ScaleIO Protection Domain for the configured storage. | false |
| readOnly | boolean | readOnly Defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts. | false |
| sslEnabled | boolean | sslEnabled Flag enable/disable SSL communication with Gateway, default false | false |
| storageMode | string | storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.<br/>Default is ThinProvisioned.<br/>*Default*: ThinProvisioned<br/> | false |
| storagePool | string | storagePool is the ScaleIO Storage Pool associated with the protection domain. | false |
| volumeName | string | volumeName is the name of a volume already created in the ScaleIO system<br/>that is associated with this volume source. | false |
### LMDeployment.spec.vllm.volumes[index].scaleIO.secretRef

secretRef references to the secret for ScaleIO user and other<br/>sensitive information. If this is not provided, Login operation will fail.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.vllm.volumes[index].secret

secret represents a secret that should populate this volume.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#secret

| Name | Type | Description | Required |
|------|------|-------------|----------|
| defaultMode | integer | defaultMode is Optional: mode bits used to set permissions on created files by default.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values<br/>for mode bits. Defaults to 0644.<br/>Directories within the path are not affected by this setting.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
| [items](#lmdeploymentspecvllmvolumesindexsecretitemsindex) | []object | items If unspecified, each key-value pair in the Data field of the referenced<br/>Secret will be projected into the volume as a file whose name is the<br/>key and content is the value. If specified, the listed keys will be<br/>projected into the specified paths, and unlisted keys will not be<br/>present. If a key is specified which is not present in the Secret,<br/>the volume setup will error unless it is marked optional. Paths must be<br/>relative and may not contain the '..' path or start with '..'. | false |
| optional | boolean | optional field specify whether the Secret or its keys must be defined | false |
| secretName | string | secretName is the name of the secret in the pod's namespace to use.<br/>More info: https://kubernetes.io/docs/concepts/storage/volumes#secret | false |
### LMDeployment.spec.vllm.volumes[index].secret.items[index]

Maps a string key to a path within a volume.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| key | string | key is the key to project. | true |
| path | string | path is the relative path of the file to map the key to.<br/>May not be an absolute path.<br/>May not contain the path element '..'.<br/>May not start with the string '..'. | true |
| mode | integer | mode is Optional: mode bits used to set permissions on this file.<br/>Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.<br/>YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.<br/>If not specified, the volume defaultMode will be used.<br/>This might be in conflict with other options that affect the file<br/>mode, like fsGroup, and the result can be other mode bits set.<br/>*Format*: int32<br/> | false |
### LMDeployment.spec.vllm.volumes[index].storageos

storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.<br/>Deprecated: StorageOS is deprecated and the in-tree storageos type is no longer supported.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| fsType | string | fsType is the filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. | false |
| readOnly | boolean | readOnly defaults to false (read/write). ReadOnly here will force<br/>the ReadOnly setting in VolumeMounts. | false |
| [secretRef](#lmdeploymentspecvllmvolumesindexstorageossecretref) | object | secretRef specifies the secret to use for obtaining the StorageOS API<br/>credentials.  If not specified, default values will be attempted. | false |
| volumeName | string | volumeName is the human-readable name of the StorageOS volume.  Volume<br/>names are only unique within a namespace. | false |
| volumeNamespace | string | volumeNamespace specifies the scope of the volume within StorageOS.  If no<br/>namespace is specified then the Pod's namespace will be used.  This allows the<br/>Kubernetes name scoping to be mirrored within StorageOS for tighter integration.<br/>Set VolumeName to any name to override the default behaviour.<br/>Set to "default" if you are not using namespaces within StorageOS.<br/>Namespaces that do not pre-exist within StorageOS will be created. | false |
### LMDeployment.spec.vllm.volumes[index].storageos.secretRef

secretRef specifies the secret to use for obtaining the StorageOS API<br/>credentials.  If not specified, default values will be attempted.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| name | string | Name of the referent.<br/>This field is effectively required, but due to backwards compatibility is<br/>allowed to be empty. Instances of this type with an empty value here are<br/>almost certainly wrong.<br/>More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names<br/>*Default*: <br/> | false |
### LMDeployment.spec.vllm.volumes[index].vsphereVolume

vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine.<br/>Deprecated: VsphereVolume is deprecated. All operations for the in-tree vsphereVolume type<br/>are redirected to the csi.vsphere.vmware.com CSI driver.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| volumePath | string | volumePath is the path that identifies vSphere volume vmdk | true |
| fsType | string | fsType is filesystem type to mount.<br/>Must be a filesystem type supported by the host operating system.<br/>Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. | false |
| storagePolicyID | string | storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName. | false |
| storagePolicyName | string | storagePolicyName is the storage Policy Based Management (SPBM) profile name. | false |
### LMDeployment.status

LMDeploymentStatus defines the observed state of Deployment

| Name | Type | Description | Required |
|------|------|-------------|----------|
| [conditions](#lmdeploymentstatusconditionsindex) | []object | Conditions represent the latest available observations of the deployment's current state | false |
| [ollamaStatus](#lmdeploymentstatusollamastatus) | object | OllamaStatus represents the status of Ollama deployment | false |
| [openwebuiStatus](#lmdeploymentstatusopenwebuistatus) | object | OpenWebUIStatus represents the status of OpenWebUI deployment | false |
| phase | string | Phase represents the current phase of the deployment | false |
| readyReplicas | integer | ReadyReplicas is the number of ready replicas<br/>*Format*: int32<br/> | false |
| [tabbyStatus](#lmdeploymentstatustabbystatus) | object | TabbyStatus represents the status of Tabby deployment | false |
| totalReplicas | integer | TotalReplicas is the total number of replicas<br/>*Format*: int32<br/> | false |
| [vllmStatus](#lmdeploymentstatusvllmstatus) | object | VLLMStatus represents the status of vLLM deployment | false |
### LMDeployment.status.conditions[index]

Condition contains details for one aspect of the current state of this API Resource.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| lastTransitionTime | string | lastTransitionTime is the last time the condition transitioned from one status to another.<br/>This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.<br/>*Format*: date-time<br/> | true |
| message | string | message is a human readable message indicating details about the transition.<br/>This may be an empty string. | true |
| reason | string | reason contains a programmatic identifier indicating the reason for the condition's last transition.<br/>Producers of specific condition types may define expected values and meanings for this field,<br/>and whether the values are considered a guaranteed API.<br/>The value should be a CamelCase string.<br/>This field may not be empty. | true |
| status | enum | status of the condition, one of True, False, Unknown.<br/>*Enum*: True, False, Unknown<br/> | true |
| type | string | type of condition in CamelCase or in foo.example.com/CamelCase. | true |
| observedGeneration | integer | observedGeneration represents the .metadata.generation that the condition was set based upon.<br/>For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date<br/>with respect to the current state of the instance.<br/>*Format*: int64<br/>*Minimum*: 0<br/> | false |
### LMDeployment.status.ollamaStatus

OllamaStatus represents the status of Ollama deployment

| Name | Type | Description | Required |
|------|------|-------------|----------|
| availableReplicas | integer | AvailableReplicas is the number of available replicas<br/>*Format*: int32<br/> | false |
| [conditions](#lmdeploymentstatusollamastatusconditionsindex) | []object | Conditions represent the latest available observations of the component's current state | false |
| readyReplicas | integer | ReadyReplicas is the number of ready replicas<br/>*Format*: int32<br/> | false |
| updatedReplicas | integer | UpdatedReplicas is the number of updated replicas<br/>*Format*: int32<br/> | false |
### LMDeployment.status.ollamaStatus.conditions[index]

Condition contains details for one aspect of the current state of this API Resource.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| lastTransitionTime | string | lastTransitionTime is the last time the condition transitioned from one status to another.<br/>This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.<br/>*Format*: date-time<br/> | true |
| message | string | message is a human readable message indicating details about the transition.<br/>This may be an empty string. | true |
| reason | string | reason contains a programmatic identifier indicating the reason for the condition's last transition.<br/>Producers of specific condition types may define expected values and meanings for this field,<br/>and whether the values are considered a guaranteed API.<br/>The value should be a CamelCase string.<br/>This field may not be empty. | true |
| status | enum | status of the condition, one of True, False, Unknown.<br/>*Enum*: True, False, Unknown<br/> | true |
| type | string | type of condition in CamelCase or in foo.example.com/CamelCase. | true |
| observedGeneration | integer | observedGeneration represents the .metadata.generation that the condition was set based upon.<br/>For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date<br/>with respect to the current state of the instance.<br/>*Format*: int64<br/>*Minimum*: 0<br/> | false |
### LMDeployment.status.openwebuiStatus

OpenWebUIStatus represents the status of OpenWebUI deployment

| Name | Type | Description | Required |
|------|------|-------------|----------|
| availableReplicas | integer | AvailableReplicas is the number of available replicas<br/>*Format*: int32<br/> | false |
| [conditions](#lmdeploymentstatusopenwebuistatusconditionsindex) | []object | Conditions represent the latest available observations of the component's current state | false |
| readyReplicas | integer | ReadyReplicas is the number of ready replicas<br/>*Format*: int32<br/> | false |
| updatedReplicas | integer | UpdatedReplicas is the number of updated replicas<br/>*Format*: int32<br/> | false |
### LMDeployment.status.openwebuiStatus.conditions[index]

Condition contains details for one aspect of the current state of this API Resource.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| lastTransitionTime | string | lastTransitionTime is the last time the condition transitioned from one status to another.<br/>This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.<br/>*Format*: date-time<br/> | true |
| message | string | message is a human readable message indicating details about the transition.<br/>This may be an empty string. | true |
| reason | string | reason contains a programmatic identifier indicating the reason for the condition's last transition.<br/>Producers of specific condition types may define expected values and meanings for this field,<br/>and whether the values are considered a guaranteed API.<br/>The value should be a CamelCase string.<br/>This field may not be empty. | true |
| status | enum | status of the condition, one of True, False, Unknown.<br/>*Enum*: True, False, Unknown<br/> | true |
| type | string | type of condition in CamelCase or in foo.example.com/CamelCase. | true |
| observedGeneration | integer | observedGeneration represents the .metadata.generation that the condition was set based upon.<br/>For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date<br/>with respect to the current state of the instance.<br/>*Format*: int64<br/>*Minimum*: 0<br/> | false |
### LMDeployment.status.tabbyStatus

TabbyStatus represents the status of Tabby deployment

| Name | Type | Description | Required |
|------|------|-------------|----------|
| availableReplicas | integer | AvailableReplicas is the number of available replicas<br/>*Format*: int32<br/> | false |
| [conditions](#lmdeploymentstatustabbystatusconditionsindex) | []object | Conditions represent the latest available observations of the component's current state | false |
| readyReplicas | integer | ReadyReplicas is the number of ready replicas<br/>*Format*: int32<br/> | false |
| updatedReplicas | integer | UpdatedReplicas is the number of updated replicas<br/>*Format*: int32<br/> | false |
### LMDeployment.status.tabbyStatus.conditions[index]

Condition contains details for one aspect of the current state of this API Resource.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| lastTransitionTime | string | lastTransitionTime is the last time the condition transitioned from one status to another.<br/>This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.<br/>*Format*: date-time<br/> | true |
| message | string | message is a human readable message indicating details about the transition.<br/>This may be an empty string. | true |
| reason | string | reason contains a programmatic identifier indicating the reason for the condition's last transition.<br/>Producers of specific condition types may define expected values and meanings for this field,<br/>and whether the values are considered a guaranteed API.<br/>The value should be a CamelCase string.<br/>This field may not be empty. | true |
| status | enum | status of the condition, one of True, False, Unknown.<br/>*Enum*: True, False, Unknown<br/> | true |
| type | string | type of condition in CamelCase or in foo.example.com/CamelCase. | true |
| observedGeneration | integer | observedGeneration represents the .metadata.generation that the condition was set based upon.<br/>For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date<br/>with respect to the current state of the instance.<br/>*Format*: int64<br/>*Minimum*: 0<br/> | false |
### LMDeployment.status.vllmStatus

VLLMStatus represents the status of vLLM deployment

| Name | Type | Description | Required |
|------|------|-------------|----------|
| availableReplicas | integer | AvailableReplicas is the number of available replicas<br/>*Format*: int32<br/> | false |
| [conditions](#lmdeploymentstatusvllmstatusconditionsindex) | []object | Conditions represent the latest available observations of the component's current state | false |
| readyReplicas | integer | ReadyReplicas is the number of ready replicas<br/>*Format*: int32<br/> | false |
| updatedReplicas | integer | UpdatedReplicas is the number of updated replicas<br/>*Format*: int32<br/> | false |
### LMDeployment.status.vllmStatus.conditions[index]

Condition contains details for one aspect of the current state of this API Resource.

| Name | Type | Description | Required |
|------|------|-------------|----------|
| lastTransitionTime | string | lastTransitionTime is the last time the condition transitioned from one status to another.<br/>This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.<br/>*Format*: date-time<br/> | true |
| message | string | message is a human readable message indicating details about the transition.<br/>This may be an empty string. | true |
| reason | string | reason contains a programmatic identifier indicating the reason for the condition's last transition.<br/>Producers of specific condition types may define expected values and meanings for this field,<br/>and whether the values are considered a guaranteed API.<br/>The value should be a CamelCase string.<br/>This field may not be empty. | true |
| status | enum | status of the condition, one of True, False, Unknown.<br/>*Enum*: True, False, Unknown<br/> | true |
| type | string | type of condition in CamelCase or in foo.example.com/CamelCase. | true |
| observedGeneration | integer | observedGeneration represents the .metadata.generation that the condition was set based upon.<br/>For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date<br/>with respect to the current state of the instance.<br/>*Format*: int64<br/>*Minimum*: 0<br/> | false |
