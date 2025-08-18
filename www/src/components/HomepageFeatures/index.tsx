import type {ReactNode} from 'react';
import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';

type FeatureItem = {
  title: string;
  Svg: React.ComponentType<React.ComponentProps<'svg'>>;
  description: ReactNode;
};

const FeatureList: FeatureItem[] = [
  {
    title: 'Kubernetes Native',
    Svg: require('@site/static/img/kubernetes-logo.svg').default,
    description: (
      <>
        Built as a native Kubernetes operator using Go and Kubebuilder. 
        Automatically manages the complete lifecycle of LLM services with 
        declarative configuration and GitOps workflows.
      </>
    ),
  },
  {
    title: 'Chat with LLM',
    Svg: require('@site/static/img/chat-llm.svg').default,
    description: (
      <>
        Interactive chat interfaces powered by OpenWebUI with real-time 
        conversations, context awareness, and seamless integration with 
        multiple LLM backends for natural AI interactions.
      </>
    ),
  },
  {
    title: 'Coding Assistants',
    Svg: require('@site/static/img/coding-assistants.svg').default,
    description: (
      <>
        Advanced coding assistance with Tabby integration, providing 
        intelligent code completion, refactoring suggestions, and 
        AI-powered development workflows for multiple programming languages.
      </>
    ),
  },
  {
    title: 'Multi-LLM Support',
    Svg: require('@site/static/img/llm-integration.svg').default,
    description: (
      <>
        Seamlessly integrates Ollama, OpenWebUI, Tabby, and custom models. 
        Supports multiple model formats, automatic scaling, and unified 
        management across different LLM backends.
      </>
    ),
  },
  {
    title: 'Advanced Observability',
    Svg: require('@site/static/img/observability.svg').default,
    description: (
      <>
        Built-in Langfuse integration for comprehensive LLM monitoring, 
        tracing, and analytics. Track requests, performance metrics, 
        costs, and user interactions in real-time.
      </>
    ),
  },
  {
    title: 'Pipeline Processing',
    Svg: require('@site/static/img/pipeline.svg').default,
    description: (
      <>
        OpenWebUI Pipelines enable custom workflows, filters, and 
        integrations. Process requests through configurable stages 
        with Python-based extensibility and automatic monitoring.
      </>
    ),
  },
  {
    title: 'Production Ready',
    Svg: require('@site/static/img/multi-component.svg').default,
    description: (
      <>
        Enterprise-grade features including Redis persistence, 
        auto-scaling, ingress management, and multi-replica support. 
        Ready for production workloads with minimal configuration.
      </>
    ),
  },
];

function Feature({title, Svg, description}: FeatureItem) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <Heading as="h3">{title}</Heading>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures(): ReactNode {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
