import React, {Component} from 'react'
import PropTypes from 'prop-types'
import ReactResizeDetector from 'react-resize-detector'
import {timeSeriesToPxSeries} from 'utils/timeSeriesTransformers'
import _ from 'lodash'

import CustomProperties from 'react-custom-properties'
import {colorsStringSchema} from 'shared/schemas'
import {ErrorHandlingWith} from 'src/shared/decorators/errors'
import InvalidData from 'src/shared/components/InvalidData'

// const validateTimeSeries = timeseries => {
//   return _.every(timeseries, r =>
//     _.every(
//       r,
//       (v, i) => (i === 0 && Date.parse(v)) || _.isNumber(v) || _.isNull(v)
//     )
//   )
// }

@ErrorHandlingWith(InvalidData)
class PxInbox extends Component {
  _isMounted = false

  constructor(props) {
    super(props)
    this.isValidData = true
    this.state = {
      height: 0,
      width: 0,
      data: [],
    }
  }

  shouldComponentUpdate(nextProps, nextState) {
    const arePropsEqual = _.isEqual(this.props, nextProps)
    const areStatesEqual = _.isEqual(this.state, nextState)
    return !arePropsEqual || !areStatesEqual
  }

  componentWillMount() {
    const {data, isInDataExplorer} = this.props
    this.parseTimeSeries(data, isInDataExplorer)
  }

  componentDidMount() {
    this._isMounted = true
  }

  componentWillUnmount() {
    this._isMounted = false
  }

  async parseTimeSeries(data) {
    const {queries} = this.props
    const arrayOfId = queries.map(item => {
      return item.queryConfig.tags.host[0]
    })
    const info = await this.getDataFromMongo(arrayOfId)
    const timeSeries = timeSeriesToPxSeries(data)
    if (this._isMounted) {
      this.setState({data: this.makeObject(info, timeSeries)})
    }
    // NEED FIX VALIDATOR!
    // this.isValidData = validateTimeSeries(
    //   _.get(this._timeSeries, 'timeSeries', [])
    // )
  }

  getDataFromMongo = async id => {
    const arrayOfInfo = []
    for (let i = 0; i < id.length; i++) {
      await fetch(`http://localhost:5000/api/device/${id[i]}`, {
        mode: 'cors',
      })
        .then(res => res.json())
        .then(info => arrayOfInfo.push(info))
        .catch(err => {
          return err
        })
    }
    return arrayOfInfo
  }

  makeObject = (info, timeSeries) => {
    const indexes = [],
      validArray = []

    timeSeries.labels.forEach((item, i) => {
      if (item === 'machine.value') {
        indexes.push(i)
      }
    })

    timeSeries.tableData.map((item, j) => {
      for (let i = 0; i < indexes.length; i++) {
        if (item[indexes[i]] === 1) {
          validArray.push({
            id: JSON.stringify(j),
            title: info[indexes[i] - 1].title,
            subtitle: info[indexes[i] - 1].subtitle,
            severity: info[indexes[i] - 1].severity,
            alerttext: info[indexes[i] - 1].alerttext,
            alertvalue: info[indexes[i] - 1].alertvalue,
            assettext: info[indexes[i] - 1].assettext,
            assetvalue: info[indexes[i] - 1].assetvalue,
            date: new Date(item[0]).toISOString(),
          })
        }
      }
    })

    return validArray
  }

  componentWillUpdate(nextProps) {
    const {data, activeQueryIndex} = this.props
    if (
      data !== nextProps.data ||
      activeQueryIndex !== nextProps.activeQueryIndex
    ) {
      this.parseTimeSeries(nextProps.data)
    }
  }

  resize = () => {
    const height = this.divElement.clientHeight
    const width = this.divElement.clientWidth
    this.setState({height})
    this.setState({width})
    this.render()
  }

  render() {
    if (!this.isValidData) {
      return <InvalidData />
    }

    const {
      // data,
      // axes,
      // title,
      // colors,
      // cellID,
      // onZoom,
      // queries,
      // hoverTime,
      // timeRange,
      // cellHeight,
      // ruleValues,
      // isBarGraph,
      isRefreshing,
      // setResolution,
      // isGraphFilled,
      // showSingleStat,
      // displayOptions,
      // staticLegend,
      // underlayCallback,
      // overrideLineColors,
      isFetchingInitially,
      // handleSetHoverTime,
    } = this.props

    // const {timeSeries} = this._timeSeries

    // If data for this graph is being fetched for the first time, show a graph-wide spinner.
    if (isFetchingInitially) {
      return <GraphSpinner />
    }

    // const options = {
    //   ...displayOptions,
    //   title,
    //   labels,
    //   rightGap: 0,
    //   yRangePad: 10,
    //   labelsKMB: true,
    //   fillGraph: true,
    //   underlayCallback,
    //   axisLabelWidth: 60,
    //   drawAxesAtZero: true,
    //   axisLineColor: '#383846',
    //   gridLineColor: '#383846',
    //   connectSeparatedPoints: true,
    // }

    // const prefix = axes ? axes.y.prefix : ''
    // const suffix = axes ? axes.y.suffix : ''
    const {width, height, data} = this.state

    const _width = `${30 * width / 100}px` // i love this type messing
    const _height = `${height}px`

    return (
      <div
        style={{height: '100%'}}
        ref={divElement => (this.divElement = divElement)}
      >
        {isRefreshing ? <GraphLoadingDots /> : null}
        <CustomProperties
          properties={{
            '--px-inbox-height': _height,
            '--px-inbox-list-width': _width,
          }}
        >
          <px-inbox-demo list-items={JSON.stringify(data)} />
        </CustomProperties>

        <ReactResizeDetector
          handleWidth={true}
          handleHeight={true}
          onResize={this.resize}
        />
      </div>
    )
  }
}

const GraphLoadingDots = () => (
  <div className="graph-panel__refreshing">
    <div />
    <div />
    <div />
  </div>
)

const GraphSpinner = () => (
  <div className="graph-fetching">
    <div className="graph-spinner" />
  </div>
)

const {array, arrayOf, bool, func, number, shape, string} = PropTypes

PxInbox.defaultProps = {
  underlayCallback: () => {},
  isGraphFilled: true,
  overrideLineColors: null,
  staticLegend: false,
}

PxInbox.propTypes = {
  cellID: string,
  axes: shape({
    y: shape({
      bounds: array,
      label: string,
    }),
    y2: shape({
      bounds: array,
      label: string,
    }),
  }),
  hoverTime: string,
  handleSetHoverTime: func,
  title: string,
  isFetchingInitially: bool,
  isRefreshing: bool,
  underlayCallback: func,
  isGraphFilled: bool,
  isBarGraph: bool,
  staticLegend: bool,
  overrideLineColors: array,
  showSingleStat: bool,
  displayOptions: shape({
    stepPlot: bool,
    stackedGraph: bool,
  }),
  activeQueryIndex: number,
  ruleValues: shape({}),
  timeRange: shape({
    lower: string.isRequired,
  }),
  isInDataExplorer: bool,
  setResolution: func,
  cellHeight: number,
  onZoom: func,
  queries: arrayOf(shape({}).isRequired).isRequired,
  data: arrayOf(shape({}).isRequired).isRequired,
  colors: colorsStringSchema,
}

export default PxInbox
