import type {ReactNode} from 'react';
import clsx from 'clsx';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Layout from '@theme/Layout';
import HomepageFeatures from '@site/src/components/HomepageFeatures';
import Heading from '@theme/Heading';

import styles from './index.module.css';

function HomepageHeader() {
  const {siteConfig} = useDocusaurusContext();
  return (
    <header className={clsx('hero hero--primary', styles.heroBanner)}>
      <div className="container">
        <Heading as="h1" className="hero__title">
          {siteConfig.title}
        </Heading>
        <p className="hero__subtitle">{siteConfig.tagline}</p>
        <p className="hero__description">
          Deploy, manage, and scale Large Language Models on Kubernetes with enterprise-grade features, 
          interactive chat interfaces, coding assistants, observability, and pipeline processing capabilities.
        </p>
        <div className={styles.buttons}>
          <Link
            className="button button--secondary button--lg"
            to="/docs/quickstart">
            Get Started - 5min ⏱️
          </Link>
          <Link
            className="button button--outline button--lg margin-left--md"
            to="/docs/overview">
            Learn More
          </Link>
        </div>
      </div>
    </header>
  );
}

export default function Home(): ReactNode {
  const {siteConfig} = useDocusaurusContext();
  return (
    <Layout
      title={`${siteConfig.title} - ${siteConfig.tagline}`}
      description="Enterprise-grade Kubernetes operator for deploying and managing Large Language Models with built-in observability, pipeline processing, interactive chat interfaces, coding assistants, and production-ready features.">
      <HomepageHeader />
      <main>
        <HomepageFeatures />
        <div className="container margin-vert--xl">
          <div className="row">
            <div className="col col--8 col--offset-2">
              <div className="text--center">
                <Heading as="h2">Why Choose LLM Operator?</Heading>
                <p className="text--lg">
                  Built for production workloads, the LLM Operator provides everything you need to run 
                  AI services at scale on Kubernetes. From interactive chat interfaces and coding assistants 
                  to enterprise deployments, get started in minutes with declarative configuration and 
                  automatic resource management.
                </p>
                <div className="margin-top--lg">
                  <Link
                    className="button button--primary button--lg"
                    to="/docs/quickstart">
                    Deploy Your First LLM
                  </Link>
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>
    </Layout>
  );
}
