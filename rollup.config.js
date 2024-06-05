import svelte from 'rollup-plugin-svelte';
import postcss from 'rollup-plugin-postcss';
import resolve from '@rollup/plugin-node-resolve';

export default ['home'].map((route) => ({
	input: `src/routes/${route}/entry.js`,
	output: {
		file: `src/routes/${route}/route.js`,
		format: 'iife',
		name: `route_${route.replaceAll('-', '_')}`,
	},
	plugins: [
		svelte({
			include: `src/routes/${route}/**/*.svelte`,
			compilerOptions: {
				// By default, the client-side compiler is used. You
				// can also use the server-side rendering compiler
				// generate: 'ssr',
				// ensure that extra attributes are added to head
				// elements for hydration (used with generate: 'ssr')
				// hydratable: true,
				// You can optionally set 'customElement' to 'true' to compile
				// your components to custom elements (aka web elements)
				// customElement: false
			}
		}),
		postcss({
			plugins: []
		}),
		resolve({
			browser: true
		}),
	]
}));

