import { productName, productSlug } from 'data/metadata'
import DocsPage from '@hashicorp/react-docs-page'
import {
  generateStaticPaths,
  generateStaticProps,
} from '@hashicorp/react-docs-page/server'

const NAV_DATA_FILE = 'data/docs-nav-data.json'
const CONTENT_DIR = 'content/docs'
const basePath = 'docs'

export default function DocsLayout(props) {
  return (
    <DocsPage
      product={{ name: productName, slug: productSlug }}
      baseRoute={basePath}
      staticProps={props}
      showVersionSelect={true}
    />
  )
}

export async function getStaticPaths() {
  return {
    fallback: 'blocking',
    paths: await generateStaticPaths({
      navDataFile: NAV_DATA_FILE,
      localContentDir: CONTENT_DIR,
      // new ----
      product: { name: productName, slug: productSlug },
      basePath: basePath,
    }),
  }
}

export async function getStaticProps({ params }) {
  try {
    return {
      props: await generateStaticProps({
        navDataFile: NAV_DATA_FILE,
        localContentDir: CONTENT_DIR,
        product: { name: productName, slug: productSlug },
        params,
        basePath: basePath,
      }),
      revalidate: 10,
    }
  } catch (err) {
    console.warn(err)
    return {
      notFound: true,
    }
  }
}
