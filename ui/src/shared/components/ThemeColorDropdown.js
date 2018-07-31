import React, {Component} from 'react'
import classnames from 'classnames'
import OnClickOutside from 'shared/components/OnClickOutside'
import {ErrorHandling} from 'src/shared/decorators/errors'
import cookie from 'react-cookies'
import CustomProperties from 'react-custom-properties'

const ThemeMenuItems = [
  {
    themeName: 'dark',
    themeDescr: 'Темная сторона',
    bradingColor: '#0F0E15',
    resizeControl: 'inline',
    headerTextColor: '#eeeeee',
    headerBackground: '#0F0E15',
    cellBgColor: '',
    formsBg: '#202028',
    formsColor: '#d4d7dd',
    formsBorder: '2px solid #383846',
    formsgBgFocus: '#202028',
    tableBorder: '#383846',
    tableHightlight: '#31313d',
    tableHeadColor: '#eeeff2',
    tableBodyColor: '#bec2cc',
    panelbg: '#292933',
    panelsolidbg: '#3c3c4b',
  },
  {
    themeName: 'light',
    themeDescr: 'Светлая тема',
    bradingColor: 'linear-gradient(#f5f7fa, #c3cfe2)',
    resizeControl: 'inline',
    headerTextColor: '#3c475f',
    headerBackground: '#73778d',
    cellBgColor: 'linear-gradient(#ebedee80, #fdfbfb80)',
    formsBg: '#d4d7dd',
    formsColor: '#202028',
    formsBorder: '2px solid #a7a8b1',
    formsBgFocus: '#5a585b',
    tableBorder: '#d1d3dd',
    tableHightlight: '#e1e3ee',
    tableHeadColor: '#202028',
    tableBodyColor: '#3c475f',
    panelbg: '#fafafc',
    panelsolidbg: '#fafafc',
  },
]

@ErrorHandling
class ThemeColorDropdown extends Component {
  constructor(props) {
    super(props)
    this.state = {
      isOpen: false,
      selectedTheme: cookie.load('selectedTheme') || 'light',
    }
  }

  handleClickOutside() {
    this.setState({isOpen: false})
  }

  handleSelection = themeName => () => {
    this.setState({isOpen: false})
    this.setState({selectedTheme: themeName})
    cookie.save('selectedTheme', themeName, {path: '/'})
  }

  toggleMenu = () => this.setState({isOpen: !this.state.isOpen})

  render() {
    const {selectedTheme} = this.state
    const {isOpen} = this.state
    const currentTheme = ThemeMenuItems.find(
      values => values.themeName === selectedTheme
    )
    return (
      <div className={classnames('autorefresh-dropdown')}>
        {/* set global styling params */}
        <CustomProperties
          global={true}
          properties={{
            // Main
            '--zsse-branding-color': currentTheme.bradingColor,
            '--zsse-resize-control': currentTheme.resizeControl,
            '--zsse-header-text-color': currentTheme.headerTextColor,
            '--zsse-header-bg': currentTheme.headerBackground,
            '--zsse-cell-bg': currentTheme.cellBgColor,
            // panels
            '--zsse-panel-bg': currentTheme.panelbg,
            '--zsse-panelsolid-bg': currentTheme.panelsolidbg,
            // forms ...
            '--zsse-form-bg': currentTheme.formsBg,
            '--zsse-form-bg-focus': currentTheme.formsBgFocus,
            '--zsse-form-color': currentTheme.formsColor,
            '--zsse-form-border': currentTheme.formsBorder,
            // tables
            '--zsse-table-border': currentTheme.tableBorder,
            '--zsse-table-highlight': currentTheme.tableHightlight,
            '--zsse-table-headercolor': currentTheme.tableHeadColor,
            '--zsee-table-body-text-color': currentTheme.tableBodyColor,
            // px....
            '--px-percent-circle-fill-color': '#308ec1',
            '--px-base-text-color': '#3c475f',
            '--px-percent-circle-text-color': '#3c475f',
            '--px-gauge-empty-color': '#b2babf',

            '--px-vis-gridlines-color': '#b2babf',
            '--px-vis-cursor-line-color': '#676a6f',

            '--px-vis-register-box': '#1be200',
            '--px-vis-register-data-value': '#434547',
            '--px-vis-register-series-name': '#4f5e7e',
            '--px-vis-register-border-color': '#ffffff00',
            '--px-vis-register-font-size': '12px',

            '--px-vis-series-color-0': '#60C1F8',

            '--px-vis-axis-color': '#b2babf',
            '--px-vis-axis-title-color': '#3c475f',
            '--px-vis-axis-inline-title-color': '#919699',
            '--px-vis-axis-inline-type-color': '#919699',
            '--px-vis-axis-inline-box-color': '#b2babf',

            '--px-tooltip-background-color': '#ffffffdc',
            '--px-tooltip-text-color': '#3c475f',
            // '--px-tooltip-light-background-color': '#f8425f',
            '--px-tooltip-light-text-color': '#3c475f',
            // '--px-tooltip-light-border-color': '#f8425f',
            '--px-vis-font-family': 'Arial',
            '--px-vis-event-line-color': '#3c475f50',
            '--px-vis-nav-brush-opacity': 0.1,
            '--px-vis-nav-brush-outline-color': '#b2babf',
          }}
        />

        {/* eof styling */}
        <div className={classnames('dropdown dropdown-120', {open: isOpen})}>
          <div
            className="btn btn-sm btn-info dropdown-toggle"
            onClick={this.toggleMenu}
          >
            <span className={classnames('icon', 'eye-open')} />
            <span className="dropdown-selected">{selectedTheme}</span>
            <span className="caret" />
          </div>
          <ul className="dropdown-menu">
            <li className="dropdown-header">Выберите тему</li>
            {ThemeMenuItems.map(item => (
              <li className="dropdown-item" key={item.themeName}>
                <a href="#" onClick={this.handleSelection(item.themeName)}>
                  {item.themeDescr}
                </a>
              </li>
            ))}
          </ul>
        </div>
      </div>
    )
  }
}

export default OnClickOutside(ThemeColorDropdown)
