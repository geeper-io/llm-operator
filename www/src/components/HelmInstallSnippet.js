import React from 'react';
import CodeBlock from '@theme/CodeBlock';
import {useDocsVersion} from '@docusaurus/plugin-content-docs/client';

export default function HelmInstallSnippet() {
    const version = useDocsVersion().version;

    return (
        <CodeBlock language="bash">{`helm install llm-operator oci://ghcr.io/geeper-io/llm-operator/llm-operator --version ${version}`}</CodeBlock>
    );
}
