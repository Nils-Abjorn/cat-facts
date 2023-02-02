# Cat facts

A vue3 scalable (and absolutly overkill) application fetching data from catfact.ninja rest API

## About the application

this application serves as a demonstration on how to configure and use vuejs in frontend as part of a large fullstack application

**it features:**

- element-plus + icons: componenent library for vue3 + icons
- vueuse + integrations: composable wrappers for vue3 (and axios integration)
- axios: used to request the API
- pinia: state management (not needed yet)
- sass: css but better (mainly used to forward variables in this demo)
- vue (of course)
- vue-router: used to fake navigation in the single page application
- typeScript: strong typing
- vitest: unit testing
- elslint and prettier for code style and syntax

**What you can see in this application yet:**

- Correct way to fetch async data and display it
- well organized files and folders
- every non specific libraries you need for a scalable vue3 app

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur) + [TypeScript Vue Plugin (Volar)](https://marketplace.visualstudio.com/items?itemName=Vue.vscode-typescript-vue-plugin).

If you don't use anything yet to make your file explorer more readable i would suggest using [Material Icon Theme](https://marketplace.visualstudio.com/items?itemName=PKief.material-icon-theme)

## Type Support for `.vue` Imports in TS

TypeScript cannot handle type information for `.vue` imports by default, so we replace the `tsc` CLI with `vue-tsc` for type checking. In editors, we need [TypeScript Vue Plugin (Volar)](https://marketplace.visualstudio.com/items?itemName=Vue.vscode-typescript-vue-plugin) to make the TypeScript language service aware of `.vue` types.

If the standalone TypeScript plugin doesn't feel fast enough to you, Volar has also implemented a [Take Over Mode](https://github.com/johnsoncodehk/volar/discussions/471#discussioncomment-1361669) that is more performant. You can enable it by the following steps:

1. Disable the built-in TypeScript Extension
   1. Run `Extensions: Show Built-in Extensions` from VSCode's command palette
   2. Find `TypeScript and JavaScript Language Features`, right click and select `Disable (Workspace)`
2. Reload the VSCode window by running `Developer: Reload Window` from the command palette.

## Customize configuration

See [Vite Configuration Reference](https://vitejs.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Type-Check, Compile and Minify for Production

```sh
npm run build
```

### Run Unit Tests with [Vitest](https://vitest.dev/)

```sh
npm run test:unit
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```
