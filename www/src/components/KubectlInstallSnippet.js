import React from 'react';
import CodeBlock from '@theme/CodeBlock';
import {useDocsVersion} from '@docusaurus/plugin-content-docs/client';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';

export default function KubectlInstallSnippet() {
    const {siteConfig} = useDocusaurusContext();
    const version = useDocsVersion().version;

    // baseUrl always ends with a slash
    const {url, baseUrl} = siteConfig;
    const fullUrl = `${url}${baseUrl}releases/${version}/install.yaml`;

    return (
        <CodeBlock language="bash">{`kubectl apply -f ${fullUrl}`}</CodeBlock>
    );
}
