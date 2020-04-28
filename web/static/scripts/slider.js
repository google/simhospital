/**
 * Copyright 2020 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

const SLIDER_NUMBER_ATTRIBUTE = 'slider_number';
const HTML_ATTRIBUTES = {
  id: 'id',
  class: 'class',
  width: 'width',
  height: 'height',
  transform: 'transform',
  points: 'points',
  placeholder: 'placeholder',
  value: 'value',
  type: 'type',
  x: 'x',
  y: 'y',
  cx: 'cx',
  cy: 'cy',
  r: 'r',
};
const HANDLE_RADIUS = 50;
const HANDLE_PADDING = 10;
const SLIDER_HEIGHT = 5;
const AXIS_HEIGHT = 20;

/**
 * Computes base slider width.
 * @return {number}
 */
function getSliderSvgWidth() {
  return window.innerWidth * 0.8;
}

/**
 * Computes base slider height.
 * @return {number}
 */
function getSliderSvgHeight() {
  return HANDLE_RADIUS * 2 + SLIDER_HEIGHT + HANDLE_PADDING + AXIS_HEIGHT;
}

/**
 * Computes range element width.
 * @return {number}
 */
function getRangeElementWidth() {
  return getSliderSvgWidth() - 2 * HANDLE_RADIUS;
}

/**
 * Computes range element height.
 * @return {number}
 */
function getRangeElementHeight() {
  return getSliderSvgHeight() - 2 * HANDLE_PADDING;
}

/**
 * Computes active slider position.
 * @return {number}
 */
function positionSliderX() {
  return d3.scaleLinear().range([0, getRangeElementWidth()]).domain([0, 400]);
}

/**
 * Appends a base slider to each range element.
 * @param {!Array<!Element>} rangeElements
 * @return {!Array<!Element>}
 */
function appendBaseSliders(rangeElements) {
  return rangeElements.map(appendBaseSlider);
}

/**
 * Appends a base slider to a single range element.
 * @param {!Element} rangeElement
 * @return {!Element}
 */
function appendBaseSlider(rangeElement) {
  const id = rangeElement.getAttribute('id');
  return d3.select(`#${id}`)
      .append('svg')
      .attr(HTML_ATTRIBUTES.width, getSliderSvgWidth())
      .attr(HTML_ATTRIBUTES.height, getSliderSvgHeight())
      .append('g')
      .attr(HTML_ATTRIBUTES.transform, `translate(${HANDLE_RADIUS}, 0)`);
}

/**
 * Initializes each base slider with corresponding, fetched start values.
 * @param {!Array<!Element>} rangeElements
 * @param {!Array<!Element>} baseSliders
 */
function initializeSliders(rangeElements, baseSliders) {
  for (const [index, rangeElement] of rangeElements.entries()) {
    d3.request(rangeElement.dataset.path)
        .get(
            data => initializeSlider(
                data.response, rangeElement, baseSliders[index], index));
  }
}

/**
 * Initializes each base slider with corresponding, fetched start values.
 * @param {number} startValue of slider
 * @param {!Element} rangeElement
 * @param {!Element} baseSlider
 * @param {number} index of slider
 */
function initializeSlider(startValue, rangeElement, baseSlider, index) {
  // the range axis
  baseSlider.append('g')
      .attr(HTML_ATTRIBUTES.class, 'range-axis')
      .attr(
          HTML_ATTRIBUTES.transform, `translate(0, ${getRangeElementHeight()})`)
      .call(d3.axisBottom(positionSliderX()));

  // the range scale
  baseSlider.append('rect')
      .attr(HTML_ATTRIBUTES.class, 'range-body')
      .attr(HTML_ATTRIBUTES.x, 0)
      .attr(HTML_ATTRIBUTES.y, getRangeElementHeight() - SLIDER_HEIGHT)
      .attr(HTML_ATTRIBUTES.width, getRangeElementWidth())
      .attr(HTML_ATTRIBUTES.height, SLIDER_HEIGHT);

  // the handle
  baseSlider.append('circle')
      .attr(HTML_ATTRIBUTES.class, 'range-dragger range-handle')
      .attr(HTML_ATTRIBUTES.cx, positionSliderX()(startValue))
      .attr(
          HTML_ATTRIBUTES.cy,
          getRangeElementHeight() - SLIDER_HEIGHT - HANDLE_RADIUS -
              HANDLE_PADDING)
      .attr(HTML_ATTRIBUTES.r, HANDLE_RADIUS)
      .attr(SLIDER_NUMBER_ATTRIBUTE, index)
      .call(d3.drag()
                .on('drag', handleDrag(baseSlider))
                .on('end', handleDragEnd(rangeElement)));

  // the label
  baseSlider.append('text')
      .attr(HTML_ATTRIBUTES.class, 'range-label')
      .attr(HTML_ATTRIBUTES.x, positionSliderX()(startValue))
      .attr(
          HTML_ATTRIBUTES.y,
          getRangeElementHeight() - SLIDER_HEIGHT - HANDLE_RADIUS -
              HANDLE_PADDING)
      .attr(HTML_ATTRIBUTES.dy, '.3em')
      .text(startValue);

  // the pointer
  baseSlider.append('polygon')
      .attr(HTML_ATTRIBUTES.class, 'range-dragger range-pointer')
      .attr(HTML_ATTRIBUTES.points, calculatePointerPoints(startValue))
      .attr(SLIDER_NUMBER_ATTRIBUTE, index)
      .call(d3.drag()
                .on('drag', handleDrag(baseSlider))
                .on('end', handleDragEnd(rangeElement)));
}

/**
 * Calculates point positions of pointer.
 * @param {number} handleValue
 * @return {string} of space-delimited point coordinates
 */
function calculatePointerPoints(handleValue) {
  const pointA = (positionSliderX()(handleValue) - (HANDLE_RADIUS / 4)) + ',' +
      (getRangeElementHeight() - SLIDER_HEIGHT - HANDLE_PADDING -
       (HANDLE_RADIUS / 10));
  const pointB = (positionSliderX()(handleValue) + (HANDLE_RADIUS / 4)) + ',' +
      (getRangeElementHeight() - SLIDER_HEIGHT - HANDLE_PADDING -
       (HANDLE_RADIUS / 10));
  const pointC = positionSliderX()(handleValue) + ',' +
      (getRangeElementHeight() - SLIDER_HEIGHT);
  return `${pointA} ${pointB} ${pointC}`;
}

/**
 * Gets percent indicated by mouse position when dragging slider.
 * @return {number}
 */
function getPercent() {
  const coordinates = d3.mouse(this);
  const xNonNegative = coordinates[0] >= 0 ? coordinates[0] : 0;
  const xLimitedToRangeWidth = xNonNegative <= getRangeElementWidth() ?
      xNonNegative :
      getRangeElementWidth();
  // find the percent represented by the mouse position
  return Math.round(positionSliderX().invert(xLimitedToRangeWidth));
}

/**
 * Computes slider positions and stores them as HTML attributes.
 *
 * Curries drag handler to bind baseSlider argument, but leaves `this` to be
 * determined at invocation time.
 *
 * @param {!Element} baseSlider
 * @return {!Function}
 */
function handleDrag(baseSlider) {
  return function() {
    const percent = getPercent.call(this);
    baseSlider.select('.range-handle')
        .attr(HTML_ATTRIBUTES.cx, positionSliderX()(percent));
    baseSlider.select('.range-label')
        .attr(HTML_ATTRIBUTES.x, positionSliderX()(percent))
        .text(percent);
    baseSlider.select('.range-pointer')
        .attr(HTML_ATTRIBUTES.points, calculatePointerPoints(percent));
  };
}

/**
 * Computes slider value and persists it to backend.
 *
 * Curries drag-end handler to bind rangeElement argument, but leaves `this` to be
 * determined at invocation time.
 * @param {!Element} rangeElement
 * @return {!Function}
 */
function handleDragEnd(rangeElement) {
  return function() {
    // find the percent represented by the mouse position
    const percent = getPercent.call(this);
    // post the new value back to the simulated hospital using path provided on
    // the on the data-path attribute of range element
    d3.request(rangeElement.dataset.path)
        .header('X-Requested-With', 'XMLHttpRequest')
        .header('Content-Type', 'application/x-www-form-urlencoded')
        .post('value=' + percent, function(data) {});
  };
}

const rangeElements = [...document.getElementsByClassName('range')];
const baseSliders = appendBaseSliders(rangeElements);
initializeSliders(rangeElements, baseSliders);
